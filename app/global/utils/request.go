package utils

import (
	"bytes"
	"context"
	"encoding/json"
	req "github.com/imroc/req/v2"
	"github.com/qiushenglei/gin-skeleton/app/configs"
	"github.com/qiushenglei/gin-skeleton/app/global/constants"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
)

// RequestB ... `requestType` 一般是 "GET" 或者 "POST"，参数放在请求Body里
func RequestB(requestType string, url string, reqData map[string]interface{}) (map[string]interface{}, error) {
	bytesData, _ := json.Marshal(reqData)
	req, _ := http.NewRequest(requestType, url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, _ := (&http.Client{}).Do(req)
	body, _ := io.ReadAll(resp.Body)

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

// HttpPost Post 请求
func HttpPost(ctx context.Context, url string, body interface{}, reply interface{}, isContinue bool) (err error) {
	// 创建一个客户端
	c := req.NewClient()

	// 如果是debug模式，打印结果到console
	if configs.AppRunMode == constants.DebugMode {
		c = c.DevMode()
	}

	// 发起post请求
	resp, err := c.R().
		SetBody(body).    // 设置request body
		SetResult(reply). // unmarshal 到 reply结构体内
		Post(url)

	err = handleResponse(resp, err)

	// 如果请求失败结束不在继续业务员，则结束goroutine(向外层抛panic)
	if isContinue == false && err != nil {
		logs.Log.Error(ctx, zap.String("url", url), zap.Any("request", body), zap.Any("response", reply))
		panic(err.Error())
	}

	// 记录日志
	logs.Log.Error(ctx, zap.String("url", url), zap.Any("request", body), zap.Any("response", reply))

	return err
}

// HttpPost Get 请求
func HttpGet(ctx context.Context, url string, body map[string]string, reply interface{}, isContinue bool) (err error) {
	// 创建一个客户端
	c := req.NewClient()

	// 如果是debug模式，打印结果到console
	if configs.AppRunMode == constants.DebugMode {
		c = c.DevMode()
	}

	// 发起post请求
	resp, err := c.R().
		SetQueryParams(body). // 设置get请求参数
		SetResult(reply).     // unmarshal 到 reply结构体内
		Get(url)

	err = handleResponse(resp, err)

	// 如果请求失败结束不在继续业务员，则结束goroutine(向外层抛panic)
	if isContinue == false && err != nil {
		logs.Log.Error(ctx, zap.String("url", url), zap.Any("request", body), zap.Any("response", reply))
		panic(err.Error())
	}

	// 记录日志
	logs.Log.Error(ctx, zap.String("url", url), zap.Any("request", body), zap.Any("response", reply))

	return err
}

func handleResponse(resp *req.Response, err error) error {

	// 请求失败，panic直接往上层抛panic
	if err != nil {
		return err
	}

	// 200~299判断业务成功
	if resp.IsSuccess() {
		return nil
	}

	// >=400 客户端错误 或 服务端错误
	if resp.IsError() {
		return err
	}

	return nil
}
