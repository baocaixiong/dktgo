package kdtgo

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Sdk struct {
	AppId      string
	Format     string
	V          string
	Secret     string
	SignMethod string
	BaseUrl    string
}

type parameters struct {
	params    map[string]string
	urlValues url.Values
}

func NewSdk(url string, appId string, secret string) *Sdk {
	sdk := &Sdk{
		AppId:      appId,
		V:          "1.0",
		SignMethod: "md5",
		Format:     "json",
		Secret:     secret,
		BaseUrl:    url,
	}

	return sdk
}

func (sdk *Sdk) SendRequest(method string, params ...map[string]string) (*Response, error) {
	p := new(parameters)
	p.urlValues = url.Values{}

	p.params = map[string]string{
		"app_id":      sdk.AppId,
		"method":      method,
		"timestamp":   NowString(),
		"format":      sdk.Format,
		"v":           sdk.V,
		"sign_method": sdk.SignMethod,
	}

	if len(params) > 0 {
		for key, value := range params[0] {
			p.params[key] = value
		}
	}

	sign := NewSign(p.params).BracketToString(sdk.Secret)
	h := md5.New()
	io.WriteString(h, sign)

	p.params["sign"] = fmt.Sprintf("%x", h.Sum(nil))

	for key, value := range p.params {
		p.urlValues.Set(key, value)
	}

	resp, err := http.PostForm(fmt.Sprintf("%s?%s", sdk.BaseUrl, p.urlValues.Encode()), p.urlValues)
	return NewResponse(resp), err
}

type Response struct {
	Resp       *http.Response
	StatusCode int
	content    []byte
	consumed   bool
}

func NewResponse(resp *http.Response) *Response {
	var r Response
	r.Resp = resp
	r.StatusCode = resp.StatusCode

	return &r
}

func (r *Response) Content() ([]byte, error) {
	if r.consumed {
		return r.content, nil
	}
	defer r.Resp.Body.Close()
	data, err := ioutil.ReadAll(r.Resp.Body)
	fmt.Println(string(data))
	if err != nil {
		return nil, err
	}
	r.consumed = true
	r.content = data
	return r.content, nil
}

func (r *Response) Json(v interface{}) error {
	b, err := r.Content()
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	return nil
}
