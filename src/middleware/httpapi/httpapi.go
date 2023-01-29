package httpapi

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func PostJson(url string, jsonStr []byte, HeaderData map[string]string) (bool, string) {

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if HeaderData != nil {
		for headerkey, headervalue := range HeaderData {
			// req.Header.Set("Content-Type", "application/json")
			req.Header.Set(headerkey, headervalue)
		}
	}
	client := &http.Client{Timeout: 10000 * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil {
		errinfo := "Post请求失败：" + err.Error()
		return false, errinfo
	}
	if resp.StatusCode != 200 {
		errinfo := "Get请求失败,状态码：" + resp.Status
		return false, errinfo
	}
	defer resp.Body.Close()

	// statuscode := resp.StatusCode
	// hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	return true, string(body)

}

func GetDefault(url string, UriData map[string]string, HeaderData map[string]string) (bool, string) {
	req, _ := http.NewRequest("GET", url, nil)
	if HeaderData != nil {
		for headerkey, headervalue := range HeaderData {
			// req.Header.Set("Content-Type", "application/json")
			req.Header.Set(headerkey, headervalue)
		}
	}
	if UriData != nil {
		q := req.URL.Query()
		for urikey, urivalue := range UriData {
			q.Add(urikey, urivalue)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{Timeout: 10000 * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil {
		errinfo := "Get请求失败" + err.Error()
		return false, errinfo
	}
	if resp.StatusCode != 200 {
		errinfo := "Get请求失败,状态码：" + resp.Status
		return false, errinfo
	}
	defer resp.Body.Close()
	// statuscode := resp.StatusCode
	// hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	return true, string(body)
}
func PutDefault(url string, UriData map[string]string, HeaderData map[string]string) (bool, string) {
	req, _ := http.NewRequest("PUT", url, nil)
	if HeaderData != nil {
		for headerkey, headervalue := range HeaderData {
			// req.Header.Set("Content-Type", "application/json")
			req.Header.Set(headerkey, headervalue)
		}
	}
	if UriData != nil {
		q := req.URL.Query()
		for urikey, urivalue := range UriData {
			q.Add(urikey, urivalue)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{Timeout: 10000 * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil {
		errinfo := "Put请求失败" + err.Error()
		return false, errinfo
	}
	if resp.StatusCode != 200 {
		errinfo := "Put请求失败,状态码：" + resp.Status
		return false, errinfo
	}
	defer resp.Body.Close()
	// statuscode := resp.StatusCode
	// hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	return true, string(body)
}

func DeleteDefault(url string, UriData map[string]string, HeaderData map[string]string) (bool, string) {
	req, _ := http.NewRequest("DELETE", url, nil)
	if HeaderData != nil {
		for headerkey, headervalue := range HeaderData {
			// req.Header.Set("Content-Type", "application/json")
			req.Header.Set(headerkey, headervalue)
		}
	}
	if UriData != nil {
		q := req.URL.Query()
		for urikey, urivalue := range UriData {
			q.Add(urikey, urivalue)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{Timeout: 10000 * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil {
		errinfo := "DELETE请求失败" + err.Error()
		return false, errinfo
	}
	if resp.StatusCode != 200 {
		errinfo := "DELETE请求失败,状态码：" + resp.Status
		return false, errinfo
	}
	defer resp.Body.Close()
	// statuscode := resp.StatusCode
	// hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	return true, string(body)
}
