package gohttp

import (
	"net/http"
	"reflect"
	"testing"
)

func TestMethod_String(t *testing.T) {
	tests := []struct {
		name string
		m    Method
		want string
	}{
		{
			"get",
			Get,
			"GET",
		},
		{
			"post",
			Post,
			"POST",
		},
		{
			"put",
			Put,
			"PUT",
		},
		{
			"delete",
			Delete,
			"DELETE",
		},
		{
			"head",
			Head,
			"HEAD",
		},
		{
			"options",
			Options,
			"OPTIONS",
		},
		{
			"patch",
			Patch,
			"PATCH",
		},
		{
			"copy",
			Copy,
			"COPY",
		},
		{
			"default",
			Method(12),
			"GET",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("Method.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		headers []map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			"case1",
			args{
				nil,
			},
			false,
		},
		{
			"case2",
			args{
				[]map[string]string{
					map[string]string{"k": "v"},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.headers...); (got == nil) != tt.wantNil {
				t.Errorf("NewClient() = %v, want nil? %v", got, tt.wantNil)
			}
		})
	}
}

func Test_sharedHeaders(t *testing.T) {
	tests := []struct {
		name string
		want map[string]string
	}{
		{
			"case1",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sharedHeaders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sharedHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Request(t *testing.T) {
	testclient := http.Client{}
	type fields struct {
		headers        map[string]string
		Client         http.Client
		requestAdapter func(req *http.Request) *http.Request
	}
	type args struct {
		m          Method
		urlString  string
		parameters []map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Requester
	}{
		{
			"case1",
			fields{
				nil,
				testclient,
				nil,
			},
			args{
				Get,
				"https://www.google.com",
				nil,
			},
			&request{
				client: &client{
					headers:        nil,
					Client:         testclient,
					requestAdapter: nil,
				},
				url:       "https://www.google.com",
				method:    Get,
				header:    nil,
				parameter: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				headers:        tt.fields.headers,
				Client:         tt.fields.Client,
				requestAdapter: tt.fields.requestAdapter,
			}
			if got := c.Request(tt.args.m, tt.args.urlString, tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Request() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_RequestAdapter(t *testing.T) {
	type fields struct {
		headers        map[string]string
		Client         http.Client
		requestAdapter func(req *http.Request) *http.Request
	}
	type args struct {
		adapter func(req *http.Request) *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Client
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				headers:        tt.fields.headers,
				Client:         tt.fields.Client,
				requestAdapter: tt.fields.requestAdapter,
			}
			if got := c.RequestAdapter(tt.args.adapter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.RequestAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_HandleRedirect(t *testing.T) {
	type fields struct {
		headers        map[string]string
		Client         http.Client
		requestAdapter func(req *http.Request) *http.Request
	}
	type args struct {
		flag bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Client
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				headers:        tt.fields.headers,
				Client:         tt.fields.Client,
				requestAdapter: tt.fields.requestAdapter,
			}
			if got := c.HandleRedirect(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.HandleRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Header(t *testing.T) {
	type fields struct {
		headers        map[string]string
		Client         http.Client
		requestAdapter func(req *http.Request) *http.Request
	}
	type args struct {
		k string
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Client
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				headers:        tt.fields.headers,
				Client:         tt.fields.Client,
				requestAdapter: tt.fields.requestAdapter,
			}
			if got := c.Header(tt.args.k, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Header() = %v, want %v", got, tt.want)
			}
		})
	}
}
