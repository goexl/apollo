package param

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
)

type Client struct {
	*Base

	Meta  string `validate:"required,url"`
	Appid string `validate:"required"`

	Logger log.Logger
	Http   *http.Client
}

func NewClient() *Client {
	return &Client{
		Base: newBase(),
	}
}
