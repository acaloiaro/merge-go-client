// This file was auto-generated by Fern from our API Definition.

package asyncpassthrough

import (
	context "context"
	fmt "fmt"
	ats "github.com/merge-api/merge-go-client/ats"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

// Asynchronously pull data from an endpoint not currently supported by Merge.
func (c *Client) Create(ctx context.Context, request *ats.DataPassthroughRequest) (*ats.AsyncPassthroughReciept, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "ats/v1/async-passthrough"

	var response *ats.AsyncPassthroughReciept
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieves data from earlier async-passthrough POST request
func (c *Client) Retrieve(ctx context.Context, asyncPassthroughReceiptId string) (*ats.AsyncPassthroughRetrieveResponse, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"ats/v1/async-passthrough/%v", asyncPassthroughReceiptId)

	var response *ats.AsyncPassthroughRetrieveResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
