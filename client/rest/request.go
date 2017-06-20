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
func (req *Request) Put(url string) (*Response, error) {
	response, err := req.request.Put(url)
	if err != nil {
		return nil, err
	}

	return NewResponse(response), nil
}

// Post executes a POST request.
func (req *Request) Post(url string) (*Response, error) {
	response, err := req.request.Post(url)
	if err != nil {
		return nil, err
	}

	return NewResponse(response), nil
}

// Delete executes a DELETE request.
func (req *Request) Delete(url string) (*Response, error) {
	response, err := req.request.Delete(url)
	if err != nil {
		return nil, err
	}

	return NewResponse(response), nil
}

// SetBody sets the request body.
func (req *Request) SetBody(body []byte) {
	req.request.SetBody(body)
}

// SetQueryParam sets a query parameter on the request.
func (req *Request) SetQueryParam(param string, value string) {
	req.request.SetQueryParam(param, value)
}

// SetFile sets a single file field name and its path for multipart upload.
func (req *Request) SetFile(param string, filePath string) {
	req.request.SetFile(param, filePath)
}

// SetFormData sets form data in the request.
func (req *Request) SetFormData(data map[string]string) {
	req.request.SetFormData(data)
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
