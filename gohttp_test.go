package gohttp

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRequest(t *testing.T) {

	type args struct {
		m          Method
		url        string
		parameters []map[string]string
	}
	tests := []struct {
		name string
		args args
		want Requester
	}{
		{
			"case2",
			args{
				Get,
				"https://www.google.com",
				nil,
			},
			&request{
				client: sharedClient.(*client),
				url:    "https://www.google.com",
				method: Get,
				header: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Request(tt.args.m, tt.args.url, tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestAdapter(t *testing.T) {
	type args struct {
		adapter func(req *http.Request) *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"case1",
			args{
				func(req *http.Request) *http.Request {
					return req
				},
			},
		},
		{
			"case2",
			args{
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RequestAdapter(tt.args.adapter)
			if reflect.ValueOf(sharedClient.(*client).requestAdapter).Pointer() != reflect.ValueOf(tt.args.adapter).Pointer() {
				t.Error("RequestAdapter not set proper")
			}
		})
	}
}
