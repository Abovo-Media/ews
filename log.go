package ews

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/Abovo-Media/go-ews/ewsxml"
)

type Logger interface {
	LogClientStart(url string, ver Version)
	LogHttpRequest(req *http.Request, body []byte)
	LogHttpResponse(resp *http.Response)
	LogResponse(resp ewsxml.Response)
}

func DefaultLogger() Logger { return new(logger) }
func NopLogger() Logger     { return new(nopLogger) }

type logger struct{}

func (l *logger) LogClientStart(url string, ver Version) {
	log.Println("EWS LogClientStart:", url, ver)
}

func (l *logger) LogHttpRequest(req *http.Request, body []byte) {
	dump, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		log.Println("Dump error:", err)
	} else {
		log.Printf("Request:\n%s%s\n----\n", dump, body)
	}
}

func (l *logger) LogHttpResponse(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Println("Dump error:", err)
	} else {
		log.Printf("LogResponse:\n%s%s\n----\n", dump)
	}
}

func (l *logger) LogResponse(resp ewsxml.Response) {
	if resp.ResponseCode == ewsxml.NoError {
		return
	}
	log.Printf("%s: %s (%s)", resp.ResponseClass, resp.MessageText, resp.ResponseCode)
}

type nopLogger struct{}

func (_ *nopLogger) LogClientStart(_ string, _ Version)       {}
func (_ *nopLogger) LogHttpRequest(_ *http.Request, _ []byte) {}
func (_ *nopLogger) LogHttpResponse(_ *http.Response)         {}
func (_ *nopLogger) LogResponse(_ ewsxml.Response)            {}
