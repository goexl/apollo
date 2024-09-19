package builder

import (
	"github.com/goexl/apollo/internal/core"
	"github.com/goexl/apollo/internal/param"
)

type Client struct {
	*base[Client]

	params *param.Client
}

func NewClient() (client *Client) {
	client = new(Client)
	client.params = param.NewClient()

	client.base = newBase(client, client.params.Base)

	return
}

func (c *Client) Meta(meta string) (client *Client) {
	c.params.Meta = meta
	client = c

	return
}

func (c *Client) Appid(appid string) (client *Client) {
	c.params.Appid = appid
	client = c

	return
}

func (c *Client) Build() *core.Client {
	return core.NewClient(c.params)
}
