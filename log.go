package ews

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/Abovo-Media/go-ews/ewsxml"
)

type Logger interface {
	LogClientStart(url string, ver Version)
	LogHttpRequest(ctx context.Context, req *http.Request, body []byte)
	LogHttpResponse(ctx context.Context, resp *http.Response)
	LogResponse(ctx context.Context, resp ewsxml.ResponseMessage)
}

func DefaultLogger() Logger { return new(logger) }

func NopLogger() Logger { return new(nopLogger) }

type logger struct{}

func (l *logger) LogClientStart(url string, ver Version) {
	log.Println("EWS LogClientStart:", url, ver)
}

func (l *logger) LogHttpRequest(ctx context.Context, req *http.Request, body []byte) {
	dump, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		log.Println("Dump error:", err)
	} else {
		log.Printf("Request:\n%s%s\n----\n", dump, body)
	}
}

func (l *logger) LogHttpResponse(ctx context.Context, resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Println("Dump error:", err)
	} else {
		log.Printf("LogResponse:\n%s%s\n----\n", dump)
	}
}

func (l *logger) LogResponse(ctx context.Context, resp ewsxml.ResponseMessage) {
	if resp.ResponseCode == ewsxml.NoError {
		return
	}
	log.Printf("%s: %s (%s)", resp.ResponseClass, resp.MessageText, resp.ResponseCode)
}

type nopLogger int

func (*nopLogger) LogClientStart(string, Version)                        {}
func (*nopLogger) LogHttpRequest(context.Context, *http.Request, []byte) {}
func (*nopLogger) LogHttpResponse(context.Context, *http.Response)       {}
func (*nopLogger) LogResponse(context.Context, ewsxml.ResponseMessage)   {}
