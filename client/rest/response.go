package rest

import (
	"fmt"

	"github.com/go-resty/resty"
)

// Response implements the Alauda REST response.
type Response struct {
	response *resty.Response
}

// Body of the response.
func (resp *Response) Body() []byte {
	return resp.response.Body()
}

// StatusCode of the response.
func (resp *Response) StatusCode() int {
	return resp.response.StatusCode()
}

// String is the response body in string.
func (resp *Response) String() string {
	return resp.response.String()
}

// CheckStatusCode returns the error message if the code indicates failure.
func (resp *Response) CheckStatusCode() error {
	code := resp.response.StatusCode()

	if code < 200 || code >= 300 {
		return fmt.Errorf("%d: %s", resp.StatusCode(), resp.String())
	}

	return nil
}

// NewResponse creates a new Alauda response.
func NewResponse(response *resty.Response) *Response {
	return &Response{
		response: response,
	}
}
