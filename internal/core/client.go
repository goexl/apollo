package core

import (
	"github.com/goexl/apollo/internal/core/internal/builder"
	"github.com/goexl/apollo/internal/param"
)

type Client struct {
	params *param.Client
}

func NewClient(params *param.Client) *Client {
	return &Client{
		params: params,
	}
}

func (c *Client) Loader() *builder.Loader {
	return builder.NewLoader(c.params)
}
