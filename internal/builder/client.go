package builder

import (
	"github.com/goexl/apollo/internal/core"
	"github.com/goexl/apollo/internal/internal/builder"
	"github.com/goexl/apollo/internal/param"
	"github.com/goexl/http"
	"github.com/goexl/log"
)

type client = builder.Base[Client]

type Client struct {
	*client

	params *param.Client
}

func NewClient() (client *Client) {
	client = new(Client)
	client.params = param.NewClient()

	client.client = builder.NewBase(client, client.params.Base)

	return
}

func (c *Client) Logger(logger log.Logger) (client *Client) {
	c.params.Logger = logger
	client = c

	return
}

func (c *Client) Http(http *http.Client) (client *Client) {
	c.params.Http = http
	client = c

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
