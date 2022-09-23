package ews

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/Abovo-Media/go-ews/ewsxml"
	"github.com/go-pogo/errors"
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
	// Argument out must be of []byte or any type that xml.Marshal
	// successfully can handle.
	Request(req *Request, out interface{}) error
}

type Client interface {
	Requester
	Log() Logger
	Url() string
	Username() string
	Do(req *Request) (*http.Response, error)
}

func NewClient(url string, ver Version, opts ...Option) (Client, error) {
	c := &client{
		log: NopLogger(),
		ver: ver,
		url: url,
		http: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: time.Second * 10,
		},
	}

	var err error
	for _, opt := range opts {
		errors.Append(&err, opt(c))
	}

	c.log.LogClientStart(url, ver)
	return c, err
}

type client struct {
	log    Logger
	http   *http.Client
	ver    Version
	url    string
	auth   [2]string
	header ewsxml.Header
}

func (c *client) Log() Logger { return c.log }

func (c *client) Url() string { return c.url }

func (c *client) Username() string { return c.auth[0] }

func (c *client) Do(req *Request) (*http.Response, error) {
	if req.head.RequestServerVersion.Version == "" {
		req.head.ServerVersion(c.ver)
	}

	body := getBuffer()
	defer releaseBuffer(body)
	if _, err := req.WriteTo(body); err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(req.ctx, http.MethodPost, c.Url(), body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if c.auth[0] != "" {
		if _, _, has := httpReq.BasicAuth(); !has {
			httpReq.SetBasicAuth(c.auth[0], c.auth[1])
		}
	}
	httpReq.Header.Set("Content-Type", "text/xml")

	c.log.LogHttpRequest(httpReq, body.Bytes())
	return c.http.Do(httpReq)
}

func (c *client) Request(req *Request, out interface{}) (err error) {
	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer errors.AppendFunc(&err, resp.Body.Close)
	c.log.LogHttpResponse(resp)

	if resp.StatusCode != http.StatusOK {
		errors.Append(&err, NewError(resp))
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	var x ewsxml.ResponseEnvelope
	if err = xml.Unmarshal(data, &x); err != nil {
		return errors.WithKind(err, UnmarshalError)
	}

	if b, ok := out.(*[]byte); ok {
		// skip unmarshalling, return as raw bytes
		*b = x.Body.Response
		return nil
	}

	return errors.WithKind(xml.Unmarshal(x.Body.Response, out), UnmarshalError)
}

var bufPool = sync.Pool{
	New: func() interface{} {
		var buf bytes.Buffer
		// len(soapStart) + len(soapBodyStart) + len(soapEnd) = 347
		buf.Grow(512)
		return &buf
	},
}

func getBuffer() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

func releaseBuffer(b *bytes.Buffer) {
	b.Reset()
	bufPool.Put(b)
}
