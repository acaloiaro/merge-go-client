// This file was auto-generated by Fern from our API Definition.

package linktoken

import (
	context "context"
	core "github.com/merge-api/merge-go-client/core"
	filestorage "github.com/merge-api/merge-go-client/filestorage"
	http "net/http"
)

type Client interface {
	Create(ctx context.Context, request *filestorage.EndUserDetailsRequest) (*filestorage.LinkToken, error)
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// Creates a link token to be used when linking a new end user.
func (c *client) Create(ctx context.Context, request *filestorage.EndUserDetailsRequest) (*filestorage.LinkToken, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/filestorage/v1/link-token"

	var response *filestorage.LinkToken
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
