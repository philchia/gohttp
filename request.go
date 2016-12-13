package gohttp

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Requester interface
type Requester interface {
	ResponseString(handler func(code int, header http.Header, body string, err error))
	ResponseData(handler func(code int, header http.Header, body []byte, err error))
	Response(handler func(resp Responsor))
}

// request is the underlying struct which implement the Requester interface
type request struct {
	client    *client
	url       string
	method    Method
	header    http.Header
	parameter map[string]string
}

// ResponseString will do the request and response with string body
func (r *request) ResponseString(handler func(code int, header http.Header, body string, err error)) {

	postValue := url.Values{}

	for k, v := range r.parameter {
		postValue.Add(k, v)
	}

	req, err := http.NewRequest(r.method.String(), r.url, strings.NewReader(postValue.Encode()))
	if err != nil {
		handler(0, nil, "", err)
		return
	}

	if r.client.requestAdapter != nil {
		req = r.client.requestAdapter(req)
	}

	resp, err := r.client.Do(req)
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
	postValue := url.Values{}

	for k, v := range r.parameter {
		postValue.Add(k, v)
	}

	req, err := http.NewRequest(r.method.String(), r.url, strings.NewReader(postValue.Encode()))
	if err != nil {
		handler(0, nil, nil, err)
		return
	}

	if r.client.requestAdapter != nil {
		req = r.client.requestAdapter(req)
	}

	resp, err := r.client.Do(req)
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

	postValue := url.Values{}

	for k, v := range r.parameter {
		postValue.Add(k, v)
	}

	req, err := http.NewRequest(r.method.String(), r.url, strings.NewReader(postValue.Encode()))
	if err != nil {
		handler(&response{err: err})
		return
	}

	if r.client.requestAdapter != nil {
		req = r.client.requestAdapter(req)
	}

	resp, err := r.client.Do(req)
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

	handler(&response{code: resp.StatusCode, body: bts, header: resp.Header, err: nil})
	return
}
