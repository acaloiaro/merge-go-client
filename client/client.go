// This file was auto-generated by Fern from our API Definition.

package client

import (
	accountingclient "github.com/merge-api/merge-go-client/accounting/client"
	atsclient "github.com/merge-api/merge-go-client/ats/client"
	core "github.com/merge-api/merge-go-client/core"
	crmclient "github.com/merge-api/merge-go-client/crm/client"
	filestorageclient "github.com/merge-api/merge-go-client/filestorage/client"
	hrisclient "github.com/merge-api/merge-go-client/hris/client"
	ticketingclient "github.com/merge-api/merge-go-client/ticketing/client"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Ats         *atsclient.Client
	Crm         *crmclient.Client
	Filestorage *filestorageclient.Client
	Hris        *hrisclient.Client
	Ticketing   *ticketingclient.Client
	Accounting  *accountingclient.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:     options.BaseURL,
		caller:      core.NewCaller(options.HTTPClient),
		header:      options.ToHeader(),
		Ats:         atsclient.NewClient(opts...),
		Crm:         crmclient.NewClient(opts...),
		Filestorage: filestorageclient.NewClient(opts...),
		Hris:        hrisclient.NewClient(opts...),
		Ticketing:   ticketingclient.NewClient(opts...),
		Accounting:  accountingclient.NewClient(opts...),
	}
}
