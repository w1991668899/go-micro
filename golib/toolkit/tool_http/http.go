package tool_http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const contenType = "application/x-www-form-urlencoded"

var (
	httpClient = &http.Client{
		Timeout: 3 * time.Second,
	}
)

func HttpPostJsonRequest(strUrl string, requestHead map[string]string, params map[string]interface{}) ([]byte, error) {
	if nil == requestHead {
		requestHead = make(map[string]string)
	}

	body := new(bytes.Buffer)
	if nil != params {
		data, _ := json.Marshal(params)
		body.Write(data)
	}

	requestHead["Content-Type"] = "application/json; charset=utf-8"
	data, err := httpRequest("POST", strUrl, requestHead, body)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func HttpPostBodyRequest(strUrl string, requestHead map[string]string, data []byte) []byte {
	if nil == requestHead {
		requestHead = make(map[string]string)
	}

	body := new(bytes.Buffer)
	if nil != data {
		body.Write(data)
	}

	requestHead["Content-Type"] = "application/json; charset=utf-8"
	data, err := httpRequest("POST", strUrl, requestHead, body)
	if err != nil {
		return nil
	} else {
		return data
	}
}

func HttpGetJsonRequest(strUrl string, requestHead map[string]string, query map[string]string) []byte {
	if nil != query {
		v := url.Values{}
		for key, value := range query {
			v.Set(key, value)
		}
		queryParams := v.Encode()
		strUrl = strUrl + "?" + queryParams
	}

	requestHead["Content-Type"] = "application/json; charset=utf-8"
	data, err := httpRequest("GET", strUrl, requestHead, nil)
	if err != nil {
		return nil
	} else {
		return data
	}
}

func HttpGetContentRequest(strUrl string, requestHead map[string]string, query map[string]string) []byte {
	if nil != query {
		v := url.Values{}
		for key, value := range query {
			v.Set(key, value)
		}
		queryParams := v.Encode()
		strUrl = strUrl + "?" + queryParams
	}

	data, err := httpRequest("GET", strUrl, requestHead, nil)
	if err != nil {
		return nil
	} else {
		return data
	}
}

func httpRequest(method, strUrl string, requestHead map[string]string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(method, strUrl, body)
	if err != nil {
		return nil, err
	}

	if requestHead != nil {
		for k, v := range requestHead {
			request.Header.Set(k, v)
		}
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	return data, err
}

//httpService Get方法
func HttpGet(requestUrl string) ([]byte, error) {
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

//httpService Post方法
func HttpPost(strUrl string, requestHead map[string]string, data []byte) ([]byte, error) {
	if nil == requestHead {
		requestHead = make(map[string]string)
	}

	body := new(bytes.Buffer)
	if nil != data {
		body.Write(data)
	}

	data, err := httpRequest("POST", strUrl, requestHead, body)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

//httpService Do方法
func HttpDo(requestUrl, jsonRequest, Method, contentType, cookie string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest(Method, requestUrl, strings.NewReader(jsonRequest))
	if err != nil {
		return nil, err
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
