// This file was auto-generated by Fern from our API Definition.

package folders

import (
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	filestorage "github.com/merge-api/merge-go-client/filestorage"
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

// Returns a list of `Folder` objects.
func (c *Client) List(ctx context.Context, request *filestorage.FoldersListRequest) (*filestorage.PaginatedFolderList, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "filestorage/v1/folders"

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
	if request.DriveId != nil {
		queryParams.Add("drive_id", fmt.Sprintf("%v", *request.DriveId))
	}
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", *request.Expand))
	}
	if request.IncludeDeletedData != nil {
		queryParams.Add("include_deleted_data", fmt.Sprintf("%v", *request.IncludeDeletedData))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.IncludeShellData != nil {
		queryParams.Add("include_shell_data", fmt.Sprintf("%v", *request.IncludeShellData))
	}
	if request.ModifiedAfter != nil {
		queryParams.Add("modified_after", fmt.Sprintf("%v", request.ModifiedAfter.Format(time.RFC3339)))
	}
	if request.ModifiedBefore != nil {
		queryParams.Add("modified_before", fmt.Sprintf("%v", request.ModifiedBefore.Format(time.RFC3339)))
	}
	if request.Name != nil {
		queryParams.Add("name", fmt.Sprintf("%v", *request.Name))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.ParentFolderId != nil {
		queryParams.Add("parent_folder_id", fmt.Sprintf("%v", *request.ParentFolderId))
	}
	if request.RemoteId != nil {
		queryParams.Add("remote_id", fmt.Sprintf("%v", *request.RemoteId))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *filestorage.PaginatedFolderList
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

// Creates a `Folder` object with the given values.
func (c *Client) Create(ctx context.Context, request *filestorage.FileStorageFolderEndpointRequest) (*filestorage.FileStorageFolderResponse, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "filestorage/v1/folders"

	queryParams := make(url.Values)
	if request.IsDebugMode != nil {
		queryParams.Add("is_debug_mode", fmt.Sprintf("%v", *request.IsDebugMode))
	}
	if request.RunAsync != nil {
		queryParams.Add("run_async", fmt.Sprintf("%v", *request.RunAsync))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *filestorage.FileStorageFolderResponse
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

// Returns a `Folder` object with the given `id`.
func (c *Client) Retrieve(ctx context.Context, id string, request *filestorage.FoldersRetrieveRequest) (*filestorage.Folder, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"filestorage/v1/folders/%v", id)

	queryParams := make(url.Values)
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", *request.Expand))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *filestorage.Folder
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

// Returns metadata for `FileStorageFolder` POSTs.
func (c *Client) MetaPostRetrieve(ctx context.Context) (*filestorage.MetaResponse, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "filestorage/v1/folders/meta/post"

	var response *filestorage.MetaResponse
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
