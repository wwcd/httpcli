package httpcli

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"time"
)

var DefaultTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

type cli struct {
	Timeout  time.Duration
	Insecure bool
}

type Option func(*cli)

func WithTimeout(d time.Duration) Option {
	return func(c *cli) { c.Timeout = d }
}

func Do(ctx context.Context, req *http.Request, opts ...Option) (*http.Response, error) {
	c := &cli{Timeout: 30 * time.Second, Insecure: true}

	for _, opt := range opts {
		opt(c)
	}

	client := http.Client{
		Timeout:   time.Duration(c.Timeout),
		Transport: DefaultTransport,
	}

	return client.Do(req.WithContext(ctx))
}

func Get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return Do(ctx, req)
}

func Post(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	return Do(ctx, req)
}

func Put(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	return Do(ctx, req)
}

func Delete(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return Do(ctx, req)
}
