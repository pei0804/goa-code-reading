// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "Sample": Sample Resource Client
//
// Command:
// $ goagen
// --design=github.com/tikasan/goa-code-reading/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-code-reading
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GETSamplePath computes a request path to the GET action of Sample.
func GETSamplePath() string {

	return fmt.Sprintf("/")
}

// Sample
func (c *Client) GETSample(ctx context.Context, path string, defaultType string, email string, enumType string, id int, integerType int, reg string, stringType string) (*http.Response, error) {
	req, err := c.NewGETSampleRequest(ctx, path, defaultType, email, enumType, id, integerType, reg, stringType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGETSampleRequest create the request corresponding to the GET action endpoint of the Sample resource.
func (c *Client) NewGETSampleRequest(ctx context.Context, path string, defaultType string, email string, enumType string, id int, integerType int, reg string, stringType string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("defaultType", defaultType)
	values.Set("email", email)
	values.Set("enumType", enumType)
	tmp3 := strconv.Itoa(id)
	values.Set("id", tmp3)
	tmp4 := strconv.Itoa(integerType)
	values.Set("integerType", tmp4)
	values.Set("reg", reg)
	values.Set("stringType", stringType)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// POSTSamplePayload is the Sample POST action payload.
type POSTSamplePayload struct {
	// デフォルト値
	DefaultType string `form:"defaultType" json:"defaultType" xml:"defaultType"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// 列挙型
	EnumType string `form:"enumType" json:"enumType" xml:"enumType"`
	// 数字（1〜10）
	IntegerType int `form:"integerType" json:"integerType" xml:"integerType"`
	// デフォルト値
	Reg string `form:"reg" json:"reg" xml:"reg"`
	// 文字（1~10文字）
	StringType string `form:"stringType" json:"stringType" xml:"stringType"`
}

// POSTSamplePath computes a request path to the POST action of Sample.
func POSTSamplePath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/%s", param0)
}

// Sample
func (c *Client) POSTSample(ctx context.Context, path string, payload *POSTSamplePayload, contentType string) (*http.Response, error) {
	req, err := c.NewPOSTSampleRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPOSTSampleRequest create the request corresponding to the POST action endpoint of the Sample resource.
func (c *Client) NewPOSTSampleRequest(ctx context.Context, path string, payload *POSTSamplePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
