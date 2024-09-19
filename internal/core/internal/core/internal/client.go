package internal

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-resty/resty/v2"
	"github.com/goexl/apollo/internal/core/internal/param"
	"github.com/goexl/exception"
	"gopkg.in/yaml.v3"
)

type Client struct {
	params *param.Loader
}

func NewClient(params *param.Loader) *Client {
	return &Client{
		params: params,
	}
}

func (c *Client) Load(target any, params *param.Loader) (err error) {
	for _, namespace := range params.Namespaces {
		err = c.load(target, namespace, params)
	}

	return
}

func (c *Client) load(target any, namespace string, params *param.Loader) (err error) {
	if content, se := c.send(namespace, params); nil != se {
		err = se
	} else if "" != content {
		ext := filepath.Ext(namespace)
		err = c.fill(ext, &content, target)
	}

	return
}

func (c *Client) fill(ext string, content *string, target any) (err error) {
	switch ext {
	case "yml", "yaml":
		err = yaml.Unmarshal([]byte(*content), target)
	case "json":
		err = json.Unmarshal([]byte(*content), target)
	case "xml":
		err = xml.Unmarshal([]byte(*content), target)
	default:
		err = json.Unmarshal([]byte(*content), target)
	}

	return
}

func (c *Client) send(namespace string, params *param.Loader) (content string, err error) {
	request := c.params.Http.R()
	form := make(map[string]string)
	if "" != params.Ip {
		form["ip"] = params.Ip
	}
	if "" != params.Label {
		form["label"] = params.Label
	}

	if "" != c.params.Key || 0 != c.params.Notification {
		content, err = c.uncached(request, namespace, &form, params)
	} else {
		content, err = c.cached(request, namespace, &form)
	}

	return
}

func (c *Client) cached(request *resty.Request, namespace string, form *map[string]string) (content string, err error) {
	request.SetFormData(*form)
	url := fmt.Sprintf("%s/configfiles/json/%s/%s/%s", c.params.Meta, c.params.Appid, c.params.Cluster, namespace)
	response := new(CachedResponse)
	if de := c.do(url, request, response); nil != de {
		err = de
	} else {
		content = response.Content
	}

	return
}

func (c *Client) uncached(
	request *resty.Request, namespace string,
	form *map[string]string, params *param.Loader,
) (content string, err error) {
	if "" != params.Key {
		(*form)["releaseKey"] = params.Key
	}
	if 0 != params.Notification {
		(*form)["messages"] = fmt.Sprintf(`{"details":{"app+default+test":%d}}`, params.Notification)
	}
	request.SetFormData(*form)

	url := fmt.Sprintf("%s/configs/%s/%s/%s", c.params.Meta, c.params.Appid, c.params.Cluster, namespace)
	response := new(Response)
	if de := c.do(url, request, response); nil != de {
		err = de
	} else {
		content = response.Configurations.Content
	}

	return
}

func (c *Client) do(url string, request *resty.Request, response any) (err error) {
	request.SetResult(response)
	if rsp, ge := request.Get(url); nil != ge {
		err = ge
	} else if http.StatusUnauthorized == rsp.StatusCode() {
		err = exception.New().Message("客户端未授权").Build()
	} else if rsp.StatusCode() == http.StatusNotFound {
		err = exception.New().Message("未找到配置项").Build()
	} else if rsp.StatusCode() == http.StatusMethodNotAllowed {
		err = exception.New().Message("接口访问方法不正确").Build()
	} else if rsp.StatusCode() == http.StatusInternalServerError {
		err = exception.New().Message("内部错误").Build()
	}

	return
}
