package conf

import (
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestParseConf(t *testing.T) {

	var str = `
[[conf]]
	method="GET"
	uri="/user/id"
	status=200
	[conf.header]
	content-type="application/json"
	cookie="a unit test"
	[conf.body]
	type="string"
	data="sample string body"
`

	f, err := ioutil.TempFile("", "gmock")
	if err != nil {
		t.Error(err)
		return
	}

	f.WriteString(str)
	f.Close()

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		wantC   Configure
		wantErr bool
	}{
		{
			name: "sample conf file with string data",
			args: args{path: f.Name()},
			wantC: Configure{[]Conf{
				{
					Method: http.MethodGet,
					URI:    "/user/id",
					Status: 200,
					Header: map[string]string{
						"content-type": "application/json",
						"cookie":       "a unit test",
					},
					Body: Body{
						Type: "string",
						Data: "sample string body",
					},
				},
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := parseConf(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("ParseConf() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestParseConf2(t *testing.T) {

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
`

	f, err := ioutil.TempFile("", "gmock")
	if err != nil {
		t.Error(err)
		return
	}

	f.WriteString(str)
	f.Close()

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		wantC   Configure
		wantErr bool
	}{
		{
			name: "sample conf file with string data",
			args: args{path: f.Name()},
			wantC: Configure{[]Conf{
				{
					Method: http.MethodPost,
					URI:    "/user/id",
					Status: 200,
					Header: map[string]string{
						"content-type": "application/json",
						"cookie":       "a unit test",
					},
					Body: Body{
						Type: "json",
						Data: map[string]interface{}{
							"result": "a json body",
						},
					},
				},
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := parseConf(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("ParseConf() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestGetConfigure(t *testing.T) {
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
`

	f, err := ioutil.TempFile("", "gmock")
	if err != nil {
		t.Error(err)
		return
	}

	f.WriteString(str)
	f.Close()

	tests := []struct {
		name    string
		wantC   Configure
		wantErr bool
	}{
		{
			name:    "return no such file error",
			wantC:   Configure{Conf: nil},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := GetConfigure()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfigure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("GetConfigure() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestGetConfigure2(t *testing.T) {
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
`

	f, err := ioutil.TempFile("", "gmock")
	if err != nil {
		t.Error(err)
		return
	}

	f.WriteString(str)
	f.Close()

	os.Setenv(ConfigPath, f.Name())
	tests := []struct {
		name    string
		wantC   Configure
		wantErr bool
	}{
		{
			name: "return new configure",
			wantC: Configure{[]Conf{
				{
					Method: http.MethodPost,
					URI:    "/user/id",
					Status: 200,
					Header: map[string]string{
						"content-type": "application/json",
						"cookie":       "a unit test",
					},
					Body: Body{
						Type: "json",
						Data: map[string]interface{}{
							"result": "a json body",
						},
					},
				}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := GetConfigure()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfigure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("GetConfigure() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
