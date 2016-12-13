package gohttp

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func Test_response_StatusCode(t *testing.T) {
	type fields struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"case1",
			fields{
				200,
			},
			200,
		},
		{
			"case2",
			fields{
				404,
			},
			404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				code: tt.fields.code,
			}
			if got := r.StatusCode(); got != tt.want {
				t.Errorf("response.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_Body(t *testing.T) {
	type fields struct {
		body []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			"case1",
			fields{
				[]byte("Hello"),
			},
			[]byte("Hello"),
		},
		{
			"case2",
			fields{
				nil,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				body: tt.fields.body,
			}
			if got := r.Body(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.Body() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_Header(t *testing.T) {
	type fields struct {
		header http.Header
	}
	tests := []struct {
		name   string
		fields fields
		want   http.Header
	}{
		{
			"case1",
			fields{
				http.Header{},
			},
			http.Header{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				header: tt.fields.header,
			}
			if got := r.Header(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.Header() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"case1",
			fields{
				errors.New("test"),
			},
			true,
		},
		{
			"case1",
			fields{
				nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				err: tt.fields.err,
			}
			if err := r.Error(); (err != nil) != tt.wantErr {
				t.Errorf("response.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
