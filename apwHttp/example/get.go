package example

import (
	"fmt"
)

type Endpoints struct {
	CurrentUserUrl   string `json:"current_user_url"`
	AuthorizationUrl string `json:"authorization_url"`
	RepositoryUrl    string `json:"repository_url"`
}

func GetEndpoint() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		// Deal with the erorr as you needed
		return nil, err
	}

	fmt.Printf("Status Code: %d", response.StatusCode())
	fmt.Printf("Status: %s", response.Status())
	fmt.Printf("Response Body: %s", response.BodyString())

	var endpoint Endpoints
	if err := response.UnmarshalJson(&endpoint); err != nil {
		// Deal with the unmarshal erorr as you needed
		return nil, err
	}

	fmt.Printf("Repository Url:%s", endpoint.RepositoryUrl)

	return &endpoint, nil
}
