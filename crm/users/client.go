// This file was auto-generated by Fern from our API Definition.

package users

import (
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	crm "github.com/merge-api/merge-go-client/crm"
	http "net/http"
	url "net/url"
	time "time"
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

// Returns a list of `User` objects.
func (c *Client) List(ctx context.Context, request *crm.UsersListRequest) (*crm.PaginatedUserList, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/crm/v1/users"

	queryParams := make(url.Values)
	if request.CreatedAfter != nil {
		queryParams.Add("created_after", fmt.Sprintf("%v", request.CreatedAfter.Format(time.RFC3339)))
	}
	if request.CreatedBefore != nil {
		queryParams.Add("created_before", fmt.Sprintf("%v", request.CreatedBefore.Format(time.RFC3339)))
	}
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.IncludeDeletedData != nil {
		queryParams.Add("include_deleted_data", fmt.Sprintf("%v", *request.IncludeDeletedData))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.IncludeRemoteFields != nil {
		queryParams.Add("include_remote_fields", fmt.Sprintf("%v", *request.IncludeRemoteFields))
	}
	if request.ModifiedAfter != nil {
		queryParams.Add("modified_after", fmt.Sprintf("%v", request.ModifiedAfter.Format(time.RFC3339)))
	}
	if request.ModifiedBefore != nil {
		queryParams.Add("modified_before", fmt.Sprintf("%v", request.ModifiedBefore.Format(time.RFC3339)))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.RemoteId != nil {
		queryParams.Add("remote_id", fmt.Sprintf("%v", *request.RemoteId))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *crm.PaginatedUserList
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

// Returns a `User` object with the given `id`.
func (c *Client) Retrieve(ctx context.Context, id string, request *crm.UsersRetrieveRequest) (*crm.User, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/crm/v1/users/%v", id)

	queryParams := make(url.Values)
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.IncludeRemoteFields != nil {
		queryParams.Add("include_remote_fields", fmt.Sprintf("%v", *request.IncludeRemoteFields))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *crm.User
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

// Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
func (c *Client) IgnoreCreate(ctx context.Context, modelId string, request *crm.IgnoreCommonModelRequest) error {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/crm/v1/users/ignore/%v", modelId)

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:     endpointURL,
			Method:  http.MethodPost,
			Headers: c.header,
			Request: request,
		},
	); err != nil {
		return err
	}
	return nil
}

// Returns a list of `RemoteFieldClass` objects.
func (c *Client) RemoteFieldClassesList(ctx context.Context, request *crm.UsersRemoteFieldClassesListRequest) (*crm.PaginatedRemoteFieldClassList, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/crm/v1/users/remote-field-classes"

	queryParams := make(url.Values)
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.IncludeDeletedData != nil {
		queryParams.Add("include_deleted_data", fmt.Sprintf("%v", *request.IncludeDeletedData))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.IncludeRemoteFields != nil {
		queryParams.Add("include_remote_fields", fmt.Sprintf("%v", *request.IncludeRemoteFields))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *crm.PaginatedRemoteFieldClassList
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
