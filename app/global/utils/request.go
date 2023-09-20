package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/imroc/req"
)

// RequestB ... `requestType` 一般是 "GET" 或者 "POST"，参数放在请求Body里
func RequestB(requestType string, url string, reqData map[string]interface{}) (map[string]interface{}, error) {
	bytesData, _ := json.Marshal(reqData)
	req, _ := http.NewRequest(requestType, url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, _ := (&http.Client{}).Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	resData := map[string]interface{}{}
	err := json.Unmarshal(body, &resData)
	if err != nil {
		//return nil, errors.New(500, "", err.Error(), nil)
		return nil, nil
	}
	return resData, nil
}

// RequestU ... `requestType` 一般是 "GET" 或者 "POST"，参数拼接在URL上
func RequestU(requestType string, url string, reqData map[string]string) (ret map[string]string, err error) {
	// 创建一个新的请求
	request, err := http.NewRequest(requestType, url, nil)
	if err != nil {
		return nil, err
	}

	// 拼接参数
	q := request.URL.Query()
	for key, val := range reqData {
		q.Add(key, val)
	}
	request.URL.RawQuery = q.Encode()

	// 发出请求
	var res *http.Response
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 读取返回的结果
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// 解编码为 map[string]string
	ret = map[string]string{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}

	// RET
	return
}

// HttpPost http 请求
func HttpPost(url string, data interface{}, reply interface{}) (err error) {
	header := req.Header{
		"Accept": "application/json",
	}
	r, err := req.Post(url, header, req.BodyJSON(data))
	if err != nil {
		return
	}

	resBody, err := r.ToBytes()
	if err != nil {
		return
	}
	defer r.Response().Body.Close()

	if err = json.Unmarshal(resBody, reply); err != nil {
		//log.ErrorLogger.Error(err.Error())
		return nil
	}

	return err
}
