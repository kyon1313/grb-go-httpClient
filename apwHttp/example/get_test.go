package example

import (
	"fmt"
	"testing"
)

func TestGetEndpoint(t *testing.T) {
	//initialization

	//executioin
	resp, err := GetEndpoint()

	//validation

	fmt.Printf("[err:%v]", err)
	fmt.Println(resp)
}
