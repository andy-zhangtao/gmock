package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gmock/conf"
)

func init() {
	var str = `
[[conf]]
	method="POST"
	uri="/user/id"
	status=200
	[conf.header]
	content-type="application/json"
	cookie="a unit test"
	[conf.body]
		type="json"
		[conf.body.data]
		result="a json body"
[[conf]]
	method="GET"
	uri="/user/id2"
	status=200
	[conf.header]
	content-type="application/xml"
	cookie="a unit test"
	[conf.body]
		type="string"
		data="a json body"
`
	f, err := ioutil.TempFile("", "gmock")
	if err != nil {
		panic(err)
	}

	f.WriteString(str)
	f.Close()

	os.Setenv(conf.ConfigPath, f.Name())

	go Run()
	time.Sleep(3 * time.Second)
}

func Test_route(t *testing.T) {

	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "test nil context",
			args: args{ctx: &gin.Context{
				Request:  nil,
				Writer:   nil,
				Params:   nil,
				Keys:     nil,
				Errors:   nil,
				Accepted: nil,
			}},
			want: nil,
		},
		{
			name: "test fill context",
			args: args{ctx: &gin.Context{
				Request:  nil,
				Writer:   nil,
				Params:   nil,
				Keys:     nil,
				Errors:   nil,
				Accepted: nil,
			}},
			want: map[string]interface{}{
				"/user/id": conf.Conf{
					Method: http.MethodPost,
					URI:    "/user/id",
					Status: 200,
					Header: map[string]string{
						"content-type": "application/json",
						"cookie":       "a unit test",
					},
					Body: conf.Body{
						Type: "json",
						Data: map[string]interface{}{
							"result": "a json body",
						},
					},
				},
				"/user/id2": conf.Conf{
					Method: http.MethodGet,
					URI:    "/user/id2",
					Status: 200,
					Header: map[string]string{
						"content-type": "application/xml",
						"cookie":       "a unit test",
					},
					Body: conf.Body{
						Type: "string",
						Data: "a json body",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != nil {
				route(tt.args.ctx)
				if len(tt.args.ctx.Keys) == 0 {
					t.Errorf("Context Keys Empty!")
				}
				if !reflect.DeepEqual(tt.args.ctx.Keys, tt.want) {
					t.Errorf("Context Error. error=%v want %v", tt.args.ctx.Keys, tt.want)
				}
			}

		})
	}
}

func TestRun(t *testing.T) {

	tests := []struct {
		name string
		want conf.Conf
		body string
	}{
		{
			name: "test one api",
			want: conf.Conf{
				Method: http.MethodPost,
				URI:    "/user/id",
				Status: 200,
				Header: map[string]string{
					"Content-Type": "application/json",
					"Cookie":       "a unit test",
				},
				Body: conf.Body{},
			},
			body: "{\"result\":\"a json body\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c, body, err := url(tt.want.URI, tt.want.Method)
			if err != nil {
				t.Errorf("Get API Error. error: %v", err)
			}

			if c.URI != tt.want.URI {
				t.Errorf("Get API Error. want: %v got: %v", tt.want.URI, c.URI)
			}

			if c.Method != tt.want.Method {
				t.Errorf("Get API Error. want: %v got: %v", tt.want.Method, c.Method)
			}

			if c.Status != tt.want.Status {
				t.Errorf("Get API Error. want: %v got: %v", tt.want.Status, c.Status)
			}

			for k, v := range tt.want.Header {
				if c.Header[k] != v {
					t.Errorf("Get API Error. want: %v:%v got: %v:%v", k, v, k, c.Header[k])
				}
			}

			if body != tt.body {
				t.Errorf("Get API Error. want: %v got: %v", tt.body, body)
			}

		})
	}
}

func TestRun2(t *testing.T) {

	tests := []struct {
		name string
		want conf.Conf
		body string
	}{
		{
			name: "test one api with string body",
			want: conf.Conf{
				Method: http.MethodGet,
				URI:    "/user/id2",
				Status: 200,
				Header: map[string]string{
					"Content-Type": "application/xml",
					"Cookie":       "a unit test",
				},
				Body: conf.Body{},
			},
			body: "a json body",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//go Run()
			//
			//time.Sleep(10 * time.Second)
			c, body, err := url(tt.want.URI, tt.want.Method)
			if err != nil {
				t.Errorf("Get API Error. error: %v", err)
			}

			if c.URI != tt.want.URI {
				t.Errorf("Get API Error. want: %v got: %v", tt.want.URI, c.URI)
			}

			if c.Method != tt.want.Method {
				t.Errorf("Get API Error. want: %v got: %v", tt.want.Method, c.Method)
			}

			if c.Status != tt.want.Status {
				t.Errorf("Get API Error. want: %v got: %v", tt.want.Status, c.Status)
			}

			for k, v := range tt.want.Header {
				if c.Header[k] != v {
					t.Errorf("Get API Error. want: %v:%v got: %v:%v", k, v, k, c.Header[k])
				}
			}

			if body != tt.body {
				t.Errorf("Get API Error. want: %v got: %v", tt.body, body)
			}

		})
	}
}

func url(uri, method string) (conf.Conf, string, error) {
	c := conf.Conf{
		URI:    uri,
		Method: method,
		Header: map[string]string{},
	}
	result := ""

	uri = fmt.Sprintf("http://127.0.0.1:8080%s", uri)

	client := &http.Client{}

	req, err := http.NewRequest(method, uri, nil)

	if err != nil {
		return c, result, err
	}

	res, err := client.Do(req)
	if err != nil {
		return c, result, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c, result, err
	}

	c.Status = res.StatusCode
	for key, val := range res.Header {
		c.Header[key] = val[0]
	}

	result = string(body)
	return c, result, nil
}
