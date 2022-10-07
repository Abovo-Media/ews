package ews

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/Azure/go-ntlmssp"
)

type Config struct {
	Url     string
	Version Version
	Timeout time.Duration
	Retries uint8
}

func (c *Config) defaults() {
	if c.Version == "" {
		c.Version = Exchange2013
	}
	if c.Timeout == 0 {
		c.Timeout = time.Second * 10
	}
	if c.Retries == 0 {
		c.Retries = 3
	}
}

type Option func(c *client) error

func WithLogger(l Logger) Option {
	return func(c *client) error {
		if l == nil {
			c.log = NopLogger()
		} else {
			c.log = l
		}
		return nil
	}
}

func WithDefaultLogger() Option { return WithLogger(DefaultLogger()) }

func WithBasicAuth(user, pass string) Option {
	return func(c *client) error {
		c.auth = [2]string{user, pass}
		return nil
	}
}

func WithTransport(t http.RoundTripper) Option {
	return func(c *client) error {
		c.http.Transport = t
		return nil
	}
}

func WithDefaultTransport(skipTls bool) Option {
	t := http.DefaultTransport
	if skipTls {
		t.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return WithTransport(t)
}

func WithNTLM(skipTls bool) Option {
	return WithTransport(new(ntlmssp.Negotiator))
}
