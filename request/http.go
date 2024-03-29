package request

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type JobRes struct {
	ApiTimeout int    `json:"apiTimeout"`
	Location   string `json:"location"`
}

func Get(url string, data interface{}) (string, error) {
	logs.Debug("请求参数：{}", data)
	dataStr := utils.ToJSON(data)
	obj := make(map[string]interface{})
	utils.FromJSON(dataStr, &obj)

	errStr, errCode := paramCheck(obj)
	if errCode != 0 {
		logs.Error("参数缺失:{}", errStr)
		return "参数缺失", errors.New(errStr)
	}

	header := GetSign("GET", url, fmt.Sprint(obj["accessKeyId"]), fmt.Sprint(obj["accessKeySecret"]))
	logs.Debug("签名结果:{}", header)
	url = fmt.Sprint(obj["host"]) + url

	paramStr := ""
	for k, v := range obj {
		if k == "host" || k == "accessKeyId" || k == "accessKeySecret" {
			continue
		}
		if IsList(v) {
			list := v.([]interface{})
			for _, s := range list {
				paramStr += fmt.Sprintf("%s=%s&", k, fmt.Sprint(s))
			}
		} else {
			switch v.(type) {
			case float64:
				paramStr += fmt.Sprintf("%s=%f&", k, v)
			case string:
				paramStr += fmt.Sprintf("%s=%s&", k, v)
			default:
				paramStr += fmt.Sprintf("%s=%v&", k, v)
			}
		}
	}
	if len(paramStr) > 0 {
		paramStr = string([]byte(paramStr)[:len(paramStr)-1]) //去除多余的&
		url = url + "?" + paramStr
	}
	logs.Debug("url:{}", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "", err
	}
	req.Header.Set("Authorization", header["Authorization"])
	req.Header.Set("Date", header["date"])
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "请求错误", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("解析错误:{}", err.Error())
		return "", err
	}

	logs.Debug("请求后的状态码:{}", resp.StatusCode)
	logs.Debug("请求后的结果:{}", string(body))
	if resp.StatusCode == 200 {
		return string(body), nil
	} else {
		return "请求错误", errors.New("请求错误" + string(body))
	}

}

func Post(url string, params interface{}) (string, error) {
	logs.Debug("请求参数：{}", params)
	dataStr := utils.ToJSON(params)
	obj := make(map[string]interface{})
	utils.FromJSON(dataStr, &obj)
	errStr, errCode := paramCheck(obj)
	if errCode != 0 {
		logs.Error("参数缺失:{}", errStr)
		return "参数缺失", errors.New(errStr)
	}

	header := GetSign("POST", url, fmt.Sprint(obj["accessKeyId"]), fmt.Sprint(obj["accessKeySecret"]))
	logs.Debug("签名结果:{}", header)
	url = fmt.Sprint(obj["host"]) + url
	logs.Debug("url:{}", url)

	resp, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:    params,
		Headers: header,
	})
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "请求错误", err
	}
	logs.Debug("返回结果:{}", resp.String())
	//需要异步调用API,除了登录相关的API外，所有不使用GET方法的API都为异步API。
	//用户调用一个异步API成功后会收到202返回码以及 Body中包含的一个轮询地址（location字段），
	//用户需要周期性的GET该轮询地址以获得API的执行结果

	if resp.StatusCode == 202 {
		res := JobRes{}
		utils.FromJSON(resp.String(), &res)
		for {
			time.Sleep(1 * time.Second)

			arr := strings.Split(res.Location, "/zstack")
			location := fmt.Sprint(obj["host"]) + "zstack" + arr[1]
			resp, err = grequests.Get(location, &grequests.RequestOptions{
				Headers: header,
			})
			if err != nil {
				logs.Error("异步调用API请求错误:{}", err.Error())
				return "请求错误", err
			}
			logs.Debug("异步请求后的状态码:{}", resp.StatusCode)
			logs.Debug("异步请求后的结果:{}", resp.String())
			if resp.StatusCode == 200 || resp.StatusCode == 503 {
				return resp.String(), nil
			} else if resp.StatusCode != 202 {
				return "请求错误", errors.New("请求错误" + resp.String())
			}
		}
	} else if resp.StatusCode == 200 {
		return resp.String(), nil
	}
	return "请求错误", errors.New("请求错误" + resp.String())
}

func Put(url string, params interface{}) (string, error) {
	logs.Debug("请求参数：{}", params)
	dataStr := utils.ToJSON(params)
	obj := make(map[string]interface{})
	utils.FromJSON(dataStr, &obj)
	errStr, errCode := paramCheck(obj)
	if errCode != 0 {
		logs.Error("参数缺失:{}", errStr)
		return "参数缺失", errors.New(errStr)
	}

	header := GetSign("PUT", url, fmt.Sprint(obj["accessKeyId"]), fmt.Sprint(obj["accessKeySecret"]))
	logs.Debug("签名结果:{}", header)
	url = fmt.Sprint(obj["host"]) + url
	logs.Debug("url:{}", url)
	resp, err := grequests.Put(url, &grequests.RequestOptions{
		JSON:    params,
		Headers: header,
	})
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "请求错误", err
	}
	logs.Debug("返回结果:{}", resp.String())
	//需要异步调用API,除了登录相关的API外，所有不使用GET方法的API都为异步API。
	//用户调用一个异步API成功后会收到202返回码以及 Body中包含的一个轮询地址（location字段），
	//用户需要周期性的GET该轮询地址以获得API的执行结果
	if resp.StatusCode == 202 {
		res := JobRes{}
		utils.FromJSON(resp.String(), &res)
		for {
			time.Sleep(1 * time.Second)
			arr := strings.Split(res.Location, "/zstack")
			location := fmt.Sprint(obj["host"]) + "zstack" + arr[1]
			resp, err = grequests.Get(location, &grequests.RequestOptions{
				Headers: header,
			})
			if err != nil {
				logs.Error("异步调用API请求错误:{}", err.Error())
				return "请求错误", err
			}
			logs.Debug("异步请求后的状态码:{}", resp.StatusCode)
			logs.Debug("异步请求后的结果:{}", resp.String())
			if resp.StatusCode == 200 || resp.StatusCode == 503 {
				return resp.String(), nil
			} else if resp.StatusCode != 202 {
				return "请求错误", errors.New("请求错误" + resp.String())
			}
		}
	} else if resp.StatusCode == 200 {
		return resp.String(), nil
	}
	return "请求错误", errors.New("请求错误" + resp.String())
}

func Delete(url string, params interface{}) (string, error) {
	logs.Debug("请求参数：{}", params)
	logs.Debug("请求参数：{}", params)
	dataStr := utils.ToJSON(params)
	obj := make(map[string]interface{})
	utils.FromJSON(dataStr, &obj)
	errStr, errCode := paramCheck(obj)
	if errCode != 0 {
		logs.Error("参数缺失:{}", errStr)
		return "参数缺失", errors.New(errStr)
	}

	header := GetSign("DELETE", url, fmt.Sprint(obj["accessKeyId"]), fmt.Sprint(obj["accessKeySecret"]))
	logs.Debug("签名结果:{}", header)
	url = fmt.Sprint(obj["host"]) + url
	logs.Debug("url:{}", url)
	resp, err := grequests.Delete(url, &grequests.RequestOptions{
		JSON:    params,
		Headers: header,
	})
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "请求错误", err
	}
	logs.Debug("返回结果:{}", resp.String())
	//需要异步调用API,除了登录相关的API外，所有不使用GET方法的API都为异步API。
	//用户调用一个异步API成功后会收到202返回码以及 Body中包含的一个轮询地址（location字段），
	//用户需要周期性的GET该轮询地址以获得API的执行结果
	if resp.StatusCode == 202 {
		res := JobRes{}
		utils.FromJSON(resp.String(), &res)
		for {
			time.Sleep(1 * time.Second)
			arr := strings.Split(res.Location, "/zstack")
			location := fmt.Sprint(obj["host"]) + "zstack" + arr[1]
			resp, err = grequests.Get(location, &grequests.RequestOptions{
				Headers: header,
			})
			if err != nil {
				logs.Error("http请求错误:{}", err.Error())
				return "请求错误", err
			}
			logs.Debug("异步请求后的状态码:{}", resp.StatusCode)
			logs.Debug("异步请求后的结果:{}", resp.String())
			if resp.StatusCode == 200 || resp.StatusCode == 503 {
				return resp.String(), nil
			} else if resp.StatusCode != 202 {
				return "请求错误", errors.New("请求错误" + resp.String())
			}
		}
	} else if resp.StatusCode == 200 {
		return resp.String(), nil
	}

	return "请求错误", errors.New("请求错误" + resp.String())
}

func DeleteUrlWithParams(url string, params interface{}) (string, error) {
	logs.Debug("请求参数：{}", params)
	dataStr := utils.ToJSON(params)
	obj := make(map[string]interface{})
	utils.FromJSON(dataStr, &obj)
	errStr, errCode := paramCheck(obj)
	if errCode != 0 {
		logs.Error("参数缺失:{}", errStr)
		return "参数缺失", errors.New(errStr)
	}
	header := GetSign("DELETE", url, fmt.Sprint(obj["accessKeyId"]), fmt.Sprint(obj["accessKeySecret"]))
	logs.Debug("签名结果:{}", header)
	url = fmt.Sprint(obj["host"]) + url
	logs.Debug("url:{}", url)

	paramStr := ""
	for k, v := range obj {
		if k == "host" || k == "accessKeyId" || k == "accessKeySecret" {
			continue
		}
		if IsList(v) {
			list := v.([]interface{})
			for _, s := range list {
				paramStr += fmt.Sprintf("%s=%s&", k, fmt.Sprint(s))
			}
		} else {
			switch v.(type) {
			case float64:
				paramStr += fmt.Sprintf("%s=%f&", k, v)
			case string:
				paramStr += fmt.Sprintf("%s=%s&", k, v)
			default:
				paramStr += fmt.Sprintf("%s=%v&", k, v)
			}
		}
	}
	if len(paramStr) > 0 {
		paramStr = string([]byte(paramStr)[:len(paramStr)-1]) //去除多余的&
		url = url + "?" + paramStr
	}

	resp, err := grequests.Delete(url, &grequests.RequestOptions{
		JSON:    params,
		Headers: header,
	})
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "请求错误", err
	}
	logs.Debug("返回结果:{}", resp.String())
	//需要异步调用API,除了登录相关的API外，所有不使用GET方法的API都为异步API。
	//用户调用一个异步API成功后会收到202返回码以及 Body中包含的一个轮询地址（location字段），
	//用户需要周期性的GET该轮询地址以获得API的执行结果
	if resp.StatusCode == 202 {
		res := JobRes{}
		utils.FromJSON(resp.String(), &res)
		for {
			time.Sleep(1 * time.Second)
			arr := strings.Split(res.Location, "/zstack")
			location := fmt.Sprint(obj["host"]) + "zstack" + arr[1]
			resp, err = grequests.Get(location, &grequests.RequestOptions{
				Headers: header,
			})
			if err != nil {
				logs.Error("http请求错误:{}", err.Error())
				return "请求错误", err
			}
			logs.Debug("异步请求后的状态码:{}", resp.StatusCode)
			logs.Debug("异步请求后的结果:{}", resp.String())
			if resp.StatusCode == 200 || resp.StatusCode == 503 {
				return resp.String(), nil
			} else if resp.StatusCode != 202 {
				return "请求错误", errors.New("请求错误" + resp.String())
			}
		}
	} else if resp.StatusCode == 200 {
		return resp.String(), nil
	}

	return "请求错误", errors.New(resp.String())
}

func paramCheck(params map[string]interface{}) (string, int) {
	_, ok := params["host"]
	if !ok {
		return "host不能为空", -1
	}
	_, ok = params["accessKeyId"]
	if !ok {
		return "accessKeyId不能为空", -1
	}
	_, ok = params["accessKeySecret"]
	if !ok {
		return "accessKeySecret不能为空", -1
	}
	return "", 0
}
