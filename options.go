package ews

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/Azure/go-ntlmssp"
)

type Option interface {
	apply(c *Client) error
}

type optionFunc func(c *Client) error

func (fn optionFunc) apply(c *Client) error { return fn(c) }

var (
	DefaultVersion          = Exchange2013
	DefaultRetries    uint8 = 3
	DefaultRetrySleep       = time.Second
)

type Config struct {
	Version    Version
	Url        string
	Username   string
	Password   string
	Retries    uint8
	RetrySleep time.Duration
}

func (conf *Config) apply(client *Client) error {
	if conf.Version != "" {
		if conf.Version == "" {
			client.Version = DefaultVersion
		} else {
			client.Version = conf.Version
		}
	}
	if conf.Url != "" {
		client.Url = conf.Url
	}
	if (conf.Username == "" && conf.Password == "") || (conf.Username != "" && conf.Password != "") {
		client.Username = conf.Username
		client.Password = conf.Password
	}
	if client.Retries == 0 {
		if conf.Retries == 0 {
			client.Retries = DefaultRetries
		} else {
			client.Retries = conf.Retries
		}
	}
	if conf.Retries != 0 {
		client.Retries = conf.Retries
	}
	if client.RetrySleep == 0 {
		if conf.RetrySleep == 0 {
			client.RetrySleep = DefaultRetrySleep
		} else {
			client.RetrySleep = conf.RetrySleep
		}
	}
	return nil
}

func WithOptions(opts ...Option) Option {
	return optionFunc(func(c *Client) error {
		return c.applyOptions(opts)
	})
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(c *Client) error {
		c.http.Timeout = t
		return nil
	})
}

func WithRetries(n uint8) Option {
	return optionFunc(func(c *Client) error {
		c.Retries = n
		return nil
	})
}

func WithRetriesAndSleep(retries uint8, sleep time.Duration) Option {
	return optionFunc(func(c *Client) error {
		c.Retries = retries
		c.RetrySleep = sleep
		return nil
	})
}

func WithBasicAuth(user, pass string) Option {
	return optionFunc(func(c *Client) error {
		c.Username = user
		c.Password = pass
		return nil
	})
}

func WithTransport(t http.RoundTripper) Option {
	return withTransport(t, false)
}

func WithDefaultTransport(skipTls bool) Option {
	return withTransport(http.DefaultTransport, skipTls)
}

func WithNTLM(skipTls bool) Option {
	return withTransport(new(ntlmssp.Negotiator), skipTls)
}

func WithSkipTLS() Option {
	return optionFunc(func(c *Client) error {
		if t, ok := c.http.Transport.(*http.Transport); ok {
			if t.TLSClientConfig == nil {
				t.TLSClientConfig = new(tls.Config)
			}
			t.TLSClientConfig.InsecureSkipVerify = true
		}
		return nil
	})
}

func withTransport(t http.RoundTripper, skipTls bool) Option {
	return optionFunc(func(c *Client) error {
		c.http.Transport = t
		if !skipTls {
			return nil
		}
		return WithSkipTLS().apply(c)
	})
}
