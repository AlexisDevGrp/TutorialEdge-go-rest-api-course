// +build e2e

package test

import (
	"fmt"
	"testing"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetComments(t *testing.T) {
	fmt.Println("Running E2E test for GetComment check endpoint")
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comments")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}
func TestPostComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug":"/", "author":"Author1", "body":"Hello test function"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}