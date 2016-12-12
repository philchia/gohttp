package gohttp

import (
	"io/ioutil"
	"net/http"
)

// Requester interface
type Requester interface {
	ResponseString(handler func(code int, header http.Header, body string, err error))
	ResponseData(handler func(code int, header http.Header, body []byte, err error))
	Response(handler func(resp Responsor))
}

// request is the underlying struct which implement the Requester interface
type request struct {
	client     *client
	err        error
	request    *http.Request
	url        string
	method     Method
	parameters map[string]interface{}
	header     http.Header
}

// ResponseString will do the request and response with string body
func (r *request) ResponseString(handler func(code int, header http.Header, body string, err error)) {
	if r.err != nil {
		handler(0, nil, "", r.err)
		return
	}
	resp, err := r.client.Do(r.request)
	if err != nil {
		handler(0, nil, "", err)
		return
	}

	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handler(0, nil, "", err)
		return
	}

	handler(resp.StatusCode, resp.Header, string(bts), err)
	return
}

// ResponseData will do the request and response with bytes body
func (r *request) ResponseData(handler func(code int, header http.Header, body []byte, err error)) {
	if r.err != nil {
		handler(0, nil, nil, r.err)
		return
	}
	resp, err := r.client.Do(r.request)
	if err != nil {
		handler(0, nil, nil, err)
		return
	}

	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handler(0, nil, nil, err)
		return
	}

	handler(resp.StatusCode, resp.Header, bts, err)
	return
}

// Response will do the request and response Response interface
func (r *request) Response(handler func(resp Responsor)) {
	if r.err != nil {
		handler(&response{err: r.err})
		return
	}
	resp, err := r.client.Do(r.request)
	if err != nil {
		handler(&response{err: err})
		return
	}

	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handler(&response{err: err})
		return
	}

	handler(&response{code: resp.StatusCode, body: bts, header: resp.Header, err: r.err})
	return
}
