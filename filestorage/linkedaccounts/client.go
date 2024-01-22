// This file was auto-generated by Fern from our API Definition.

package linkedaccounts

import (
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	filestorage "github.com/merge-api/merge-go-client/filestorage"
	http "net/http"
	url "net/url"
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

// List linked accounts for your organization.
func (c *Client) List(ctx context.Context, request *filestorage.LinkedAccountsListRequest) (*filestorage.PaginatedAccountDetailsAndActionsList, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/filestorage/v1/linked-accounts"

	queryParams := make(url.Values)
	if request.Category != nil {
		queryParams.Add("category", fmt.Sprintf("%v", *request.Category))
	}
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.EndUserEmailAddress != nil {
		queryParams.Add("end_user_email_address", fmt.Sprintf("%v", *request.EndUserEmailAddress))
	}
	if request.EndUserOrganizationName != nil {
		queryParams.Add("end_user_organization_name", fmt.Sprintf("%v", *request.EndUserOrganizationName))
	}
	if request.EndUserOriginId != nil {
		queryParams.Add("end_user_origin_id", fmt.Sprintf("%v", *request.EndUserOriginId))
	}
	if request.EndUserOriginIds != nil {
		queryParams.Add("end_user_origin_ids", fmt.Sprintf("%v", *request.EndUserOriginIds))
	}
	if request.Id != nil {
		queryParams.Add("id", fmt.Sprintf("%v", *request.Id))
	}
	if request.Ids != nil {
		queryParams.Add("ids", fmt.Sprintf("%v", *request.Ids))
	}
	if request.IncludeDuplicates != nil {
		queryParams.Add("include_duplicates", fmt.Sprintf("%v", *request.IncludeDuplicates))
	}
	if request.IntegrationName != nil {
		queryParams.Add("integration_name", fmt.Sprintf("%v", *request.IntegrationName))
	}
	if request.IsTestAccount != nil {
		queryParams.Add("is_test_account", fmt.Sprintf("%v", *request.IsTestAccount))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.Status != nil {
		queryParams.Add("status", fmt.Sprintf("%v", *request.Status))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *filestorage.PaginatedAccountDetailsAndActionsList
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
