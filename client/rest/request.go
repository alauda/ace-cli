package rest

import (
	"fmt"

	"github.com/go-resty/resty"
)

// Request implements the Alauda REST request.
type Request struct {
	request *resty.Request
}

// Get executes a GET request.
func (req *Request) Get(url string) (*Response, error) {
	response, err := req.request.Get(url)
	if err != nil {
		return nil, err
	}

	return NewResponse(response), nil
}

// Put executes a PUT request.
func (req *Request) Put(url string, body []byte) (*Response, error) {
	return nil, nil
}

// Post executes a POST request.
func (req *Request) Post(url string, body []byte) (*Response, error) {
	req.request.SetBody(body)

	response, err := req.request.Post(url)
	if err != nil {
		return nil, err
	}

	return NewResponse(response), nil
}

// SetQueryParam sets a query parameter on the request.
func (req *Request) SetQueryParam(param string, value string) {
	req.request.SetQueryParam(param, value)
}

// NewRequest creates a base Request object.
func NewRequest(token string) *Request {
	restyReq := resty.R()
	restyReq.SetHeader("Content-Type", "application/json")

	if token != "" {
		restyReq.SetHeader("Authorization", fmt.Sprintf("Token %s", token))
	}

	return &Request{
		request: restyReq,
	}
}
