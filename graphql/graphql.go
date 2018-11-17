package graphql

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	"google.golang.org/appengine/urlfetch"
)

type Client struct {
	client      *graphql.Client
	accessToken string
}

func NewAppEngineClient(shop string, token string, ctx context.Context) *Client {
	httpClient := urlfetch.Client(ctx)
	return NewClient(shop, token, graphql.WithHTTPClient(httpClient))
}

func NewClient(shop string, token string, options graphql.ClientOption) *Client {
	endpoint := fmt.Sprintf("https://%s/admin/api/graphql.json", shop)
	return &Client{
		accessToken: token,
		client:      graphql.NewClient(endpoint, options),
	}
}

func (c *Client) RunRequest(ctx context.Context, q string, resp interface{}) error {
	req := graphql.NewRequest(q)
	req.Header.Set("X-Shopify-Access-Token", c.accessToken)
	return c.client.Run(ctx, req, &resp)
}
