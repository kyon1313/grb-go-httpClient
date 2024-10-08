package apw_http

import (
	"net/http"
	"sync"
)

type Client interface {
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, headers http.Header, body any) (*Response, error)
	Put(url string, headers http.Header, body any) (*Response, error)
	Patch(url string, headers http.Header, body any) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
}

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

func (c *httpClient) Get(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body any) (*Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body any) (*Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body any) (*Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
