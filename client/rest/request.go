package rest

import "github.com/go-resty/resty"

// Request implements the Alauda REST request.
type Request struct {
	request *resty.Request
}

// Get executes a GET request.
func (req *Request) Get(url string) (*Response, error) {
	return nil, nil
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

// NewRequest creates a base Request object.
func NewRequest() *Request {
	restyReq := resty.R()
	restyReq.SetHeader("Content-Type", "application/json")

	return &Request{
		request: restyReq,
	}
}
