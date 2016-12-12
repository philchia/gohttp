package gohttp

import (
	"net/http"
)

// Responsor interface is the payload of response
type Responsor interface {
	StatusCode() int
	Body() []byte
	Header() http.Header
	Error() error
}

// response is the underly struct which implement the Responsor interface
type response struct {
	code   int
	body   []byte
	header http.Header
	err    error
}

// StatusCode return the response's status code
func (r *response) StatusCode() int {
	return r.code
}

// Body return the response's bytes body
func (r *response) Body() []byte {
	return r.body
}

// Header return the response's header
func (r *response) Header() http.Header {
	return r.header
}

// Error return any error occured
func (r *response) Error() error {
	return r.err
}
