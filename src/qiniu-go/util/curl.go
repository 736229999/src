package util

import (
	"strings"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func CurlGetReturnString(url string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)
	return string(body)
}


/**
	不带参数curlget请求
 */
func CurlGetWithNoParams(uri string)(map[string]interface{},error)  {
	resp,err := http.Get(uri)
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	m := map[string]interface{}{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func CurlGetReturnMapString(uri string) (map[string]string,error)  {
	resp,err := http.Get(uri)
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	m := map[string]string{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}


/**
	带参数curl的get请求
 */
func CurlGetWithParams(uri string, param url.Values) (map[string]interface{}, error) {
	if param != nil && len(param) > 0 {
		values := param.Encode()
		uri += "?" + values
	}
	resp, err := http.Get(uri)
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	m := map[string]interface{}{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

/**
	curl的post请求
	@params表示传入的参数，eg："name=jerry&age=10"
 */
func CurlPost(url ,params string ) string{
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(params))
		//strings.NewReader("name=cjb"))
	CheckError(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)

	return string(body)
}
