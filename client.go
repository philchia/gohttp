package gohttp

import "net/http"

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
	// Options represent http OPTION method
	Options
	// Patch represent http PATCH method
	Patch
	// Copy represent http COPY method
	Copy
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
	case Options:
		return "OPTIONS"
	case Patch:
		return "PATCH"
	case Copy:
		return "COPY"
	default:
		return "GET"
	}
}

// Client interface
type Client interface {
	Request(m Method, url string, parameters ...map[string]string) Requester
	RequestAdapter(adapter func(req *http.Request) *http.Request) Client
	HandleRedirect(flag bool) Client
	Header(k, v string) Client
}

// NewClient create a client with custom headers
func NewClient(headers ...map[string]string) Client {
	var header map[string]string
	if len(headers) > 0 {
		header = headers[0]
	} else {
		header = sharedHeaders()
	}

	return &client{
		headers: header,
		Client:  http.Client{},
	}
}

func sharedHeaders() map[string]string {
	return nil
}

type client struct {
	headers map[string]string
	http.Client
	requestAdapter func(req *http.Request) *http.Request
}

// Request will make a request use the given client
func (c *client) Request(m Method, urlString string, parameters ...map[string]string) Requester {

	r := request{
		client: c,
		url:    urlString,
		method: m,
	}
	if len(parameters) > 0 {
		r.parameter = parameters[0]
	}
	return &r
}

func (c *client) RequestAdapter(adapter func(req *http.Request) *http.Request) Client {
	c.requestAdapter = adapter
	return c
}

func (c *client) HandleRedirect(flag bool) Client {
	if flag {
		c.Client.CheckRedirect = nil
	} else {
		c.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return c
}

func (c *client) Header(k, v string) Client {
	c.headers[k] = v
	return c
}
