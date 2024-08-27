package apw_http

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMaxConnection             = 5
	defaultResponseTimeoutDuration   = time.Second * 5
	defaultConnectionTimeoutDuration = time.Second * 1
)

// here i put all the basic things we need for request
// the header in the parameter is the header set per request
func (c *httpClient) do(httpMethod, url string, headers http.Header, body any) (*Response, error) {

	fullHeaders := c.getRequestHeader(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create new request")
	}

	request.Header = fullHeaders

	//when the first request comes i, he will make the new client and set all the configurations
	//after that. all the coming request will use the same client created by the first request
	client := c.getHttpClient()

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	finalResponse := &Response{
		status:     resp.Status,
		statusCode: resp.StatusCode,
		headers:    resp.Header,
		body:       respBody,
	}

	return finalResponse, nil
}

func (c *httpClient) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {

		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				//this should be configure based on the traffic in your application
				//the ammount of connection we want to have as idle in our connection\
				//idle meaning that you have connections in there doing nothing but waiting for an incoming request
				//and then you can use that connection to perform a request without opening a brand new connection
				MaxIdleConnsPerHost: c.getMaxIdleConnections(),
				//this is the amount of time we;re gonna wait for the response to comeback
				ResponseHeaderTimeout: c.getResponseTimeout(),

				//this is how long we're gonna wait for a new connection
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	//this is the default client

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnection > 0 {
		return c.builder.maxIdleConnection
	}

	return defaultMaxConnection
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}

	if c.builder.disableTimeout {
		return 0
	}

	return defaultResponseTimeoutDuration
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}

	if c.builder.disableTimeout {
		return 0
	}

	return defaultConnectionTimeoutDuration
}

func (c *httpClient) getRequestHeader(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}

func (v *httpClient) getRequestBody(contentType string, body any) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		//we ahve a default json body
		return json.Marshal(body)
	}

}
