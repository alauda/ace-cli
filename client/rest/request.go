package rest

import (
	"fmt"

	"github.com/go-resty/resty"
)

// Request implements the Alauda REST request.
type Request struct {
	*resty.Request
}

// Get executes a GET request.
func (req *Request) Get(url string) (*Response, error) {
	response, err := req.Request.Get(url)
	if err != nil {
		return nil, err
	}

	return NewResponse(response), nil
}

// Put executes a PUT request.
func (req *Request) Put(url string) (*Response, error) {
	response, err := req.Request.Put(url)
	if err != nil {
		return nil, err
	}

	return NewResponse(response), nil
}

// Post executes a POST request.
func (req *Request) Post(url string) (*Response, error) {
	response, err := req.Request.Post(url)
	if err != nil {
		return nil, err
	}

	return NewResponse(response), nil
}

// Delete executes a DELETE request.
func (req *Request) Delete(url string) (*Response, error) {
	response, err := req.Request.Delete(url)
	if err != nil {
		return nil, err
	}

	return NewResponse(response), nil
}

// NewRequest creates a base Request object.
func NewRequest(token string) *Request {
	restyReq := resty.R()
	restyReq.SetHeader("Content-Type", "application/json")

	if token != "" {
		restyReq.SetHeader("Authorization", fmt.Sprintf("Token %s", token))
	}

	return &Request{
		restyReq,
	}
}
