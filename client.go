package ews

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Abovo-Media/go-ews/ewsxml"
	"github.com/go-pogo/errors"
	"github.com/go-pogo/writing"
)

type Version = ewsxml.Version

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	Exchange2010     Version = "Exchange2010"
	Exchange2013     Version = "Exchange2013"
	Exchange2013_SP1 Version = "Exchange2013_SP1"

	RequestError   errors.Kind = "request error"
	UnmarshalError errors.Kind = "unmarshal error"
)

type Requester interface {
	// Request
	// Argument out must be of type []byte or any type that xml.Marshal
	// successfully can handle.
	Request(req *Request, out interface{}) error
}

type Client struct {
	Config

	log  Logger
	http *http.Client
}

func NewClient(url string, ver Version, opts ...Option) (*Client, error) {
	c := &Client{
		log: NopLogger(),
		http: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: time.Second * 10,
		},
		Config: Config{
			Url:        url,
			Version:    ver,
			Retries:    DefaultRetries,
			RetrySleep: DefaultRetrySleep,
		},
	}
	if err := c.applyOptions(opts); err != nil {
		return nil, err
	}

	c.log.LogClientStart(c.Url, c.Version)
	return c, nil
}

func (c *Client) applyOptions(opts []Option) error {
	var err error
	for _, opt := range opts {
		errors.Append(&err, opt.apply(c))
	}
	return err
}

var bufPool = writing.NewBytesBufferPool(512)

func (c *Client) Do(req *Request) (*http.Response, error) {
	if req.head.ServerVersion() == "" {
		req.head.WithServerVersion(c.Version)
	}

	body := bufPool.Get()
	defer bufPool.Put(body)
	if _, err := req.WriteTo(body); err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(req.ctx, http.MethodPost, c.Url, body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if c.Username != "" && c.Password != "" {
		if _, _, has := httpReq.BasicAuth(); !has {
			httpReq.SetBasicAuth(c.Username, c.Password)
		}
	}

	httpReq.Header.Set("Content-Type", "text/xml")
	c.log.LogHttpRequest(req.ctx, httpReq, body.Bytes())

	var httpResp *http.Response
	var attempt uint8
	sleep := c.RetrySleep

	for {
		attempt++
		httpResp, err = c.do(context.WithValue(req.ctx, attemptKey{}, attempt), httpReq)
		if err == nil && httpResp.StatusCode == http.StatusOK {
			return httpResp, nil
		}
		if attempt >= c.Retries {
			break
		}

		time.Sleep(sleep)
		sleep += sleep
	}
	return httpResp, err
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	// todo: record metrics
	resp, err := c.http.Do(req)

	if err != nil {
		return nil, errors.WithKind(err, RequestError)
	}

	c.log.LogHttpResponse(ctx, resp)
	return resp, nil
}

func (c *Client) Request(req *Request, out interface{}) (err error) {
	httpResp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer errors.AppendFunc(&err, httpResp.Body.Close)
	if httpResp.StatusCode != http.StatusOK {
		errors.Append(&err, NewError(httpResp))
		return err
	}

	data, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	var resp ewsxml.ResponseEnvelope
	if err = xml.Unmarshal(data, &resp); err != nil {
		return errors.WithKind(err, UnmarshalError)
	}

	if b, ok := out.(*[]byte); ok {
		// skip unmarshalling, return as raw bytes
		*b = resp.Body.Response
		return nil
	}

	// unmarshal raw response into the expected result
	if err = xml.Unmarshal(resp.Body.Response, out); err != nil {
		return errors.WithKind(err, UnmarshalError)
	}
	if resp, ok := out.(ewsxml.Response); ok {
		c.log.LogResponse(req.ctx, *resp.Response())
		if err = newResponseError(resp.Response()); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

type ResponseError struct {
	Response ewsxml.Response
}

func newResponseError(resp ewsxml.Response) error {
	if resp.Response().ResponseClass != ewsxml.ResponseClass_Error {
		return nil
	}
	return &ResponseError{Response: resp}
}

func (re *ResponseError) Error() string {
	r := re.Response.Response()
	return fmt.Sprintf("response error: %s: %s", r.ResponseCode, r.MessageText)
}
