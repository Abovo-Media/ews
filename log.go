package ews

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/Abovo-Media/go-ews/ewsxml"
)

type Logger interface {
	NewClient(conf Config)
	HttpRequest(ctx context.Context, req *http.Request, body []byte)
	HttpResponse(ctx context.Context, resp *http.Response)
	Response(ctx context.Context, resp ewsxml.ResponseMessage)
}

func WithLogger(l Logger) Option {
	return optionFunc(func(c *Client) error {
		if l == nil {
			c.log = NopLogger()
		} else {
			c.log = l
		}
		return nil
	})
}

type DefaultLogger struct {
	Log *log.Logger
}

func (l *DefaultLogger) log() *log.Logger {
	if l.Log == nil {
		l.Log = log.Default()
	}
	return l.Log
}

func (l *DefaultLogger) NewClient(conf Config) {
	l.log().Println("EWS NewClient:", conf.Url, conf.Version)
}

func (l *DefaultLogger) HttpRequest(ctx context.Context, req *http.Request, body []byte) {
	dump, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		l.log().Println("Dump error:", err)
	} else {
		l.log().Printf("Request:\n%s%s\n----\n", dump, body)
	}
}

func (l *DefaultLogger) HttpResponse(ctx context.Context, resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		l.log().Println("Dump error:", err)
	} else {
		l.log().Printf("Response:\n%s%s\n----\n", dump)
	}
}

func (l *DefaultLogger) Response(ctx context.Context, resp ewsxml.ResponseMessage) {
	if resp.ResponseCode == ewsxml.NoError {
		return
	}
	l.log().Printf("%s: %s (%s)", resp.ResponseClass, resp.MessageText, resp.ResponseCode)
}

func NopLogger() Logger { return new(nopLogger) }

type nopLogger struct{}

func (*nopLogger) NewClient(Config)                                   {}
func (*nopLogger) HttpRequest(context.Context, *http.Request, []byte) {}
func (*nopLogger) HttpResponse(context.Context, *http.Response)       {}
func (*nopLogger) Response(context.Context, ewsxml.ResponseMessage)   {}
