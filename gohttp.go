package gohttp

import (
	"net/http"
	"net/http/cookiejar"
)

var sharedClient Client

func init() {
	cookieJar, _ := cookiejar.New(nil)
	sharedClient = &client{
		headers: map[string]string{"User-Agent": "gohttp"},
		Client: http.Client{
			Jar: cookieJar,
		},
	}
}

// Request will make a request use the shared client
func Request(m Method, url string, parameters ...map[string]interface{}) Requester {
	return sharedClient.Request(m, url, parameters...)
}
