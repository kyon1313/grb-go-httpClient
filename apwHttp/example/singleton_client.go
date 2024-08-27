package example

import (
	"time"

	apw_http "github.com/kyon1313/grb-go-httpClient/apwHttp/go-httpClient"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() apw_http.Client {
	client := apw_http.NewClientBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponsetimeout(3 * time.Second).
		Build()

	return client
}
