package gohttp

import (
	"net/http"
	"net/url"
	"strings"
)

// Method represent http method
type Method int8

const (
	// Get represnet http GET method
	Get Method = iota
	// Post represnet http POST method
	Post
	// Put represnet http PUT method
	Put
	// Delete represnet http DELETE method
	Delete
	// Head represent http HEAD method
	Head
	// Option represent http OPTION method
	Option
)

// String will return http method string base on the Method type
func (m Method) String() string {
	switch m {
	case Get:
		return "GET"
	case Post:
		return "POST"
	case Put:
		return "PUT"
	case Delete:
		return "DELETE"
	case Head:
		return "HEAD"
	case Option:
		return "OPTION"
	default:
		return "GET"
	}
}

// Client interface
type Client interface {
	Request(m Method, url string, parameters ...map[string]string) Requester
	RequestAdapter(adapter func(req *http.Request) *http.Request)
}

type client struct {
	headers map[string]string
	http.Client
	requestAdapter func(req *http.Request) *http.Request
}

// Request will make a request use the given client
func (c *client) Request(m Method, urlString string, parameters ...map[string]string) Requester {

	var parameter map[string]string
	if len(parameters) > 0 {
		parameter = parameters[0]
	}

	postValue := url.Values{}

	for k, v := range parameter {
		postValue.Add(k, v)
	}

	req, err := http.NewRequest(m.String(), urlString, strings.NewReader(postValue.Encode()))
	if c.requestAdapter != nil {
		req = c.requestAdapter(req)
	}

	r := request{
		err:     err,
		client:  c,
		request: req,
		url:     urlString,
		method:  m,
	}
	return &r
}

func (c *client) RequestAdapter(adapter func(req *http.Request) *http.Request) {
	c.requestAdapter = adapter
}
