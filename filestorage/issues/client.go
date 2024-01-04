// This file was auto-generated by Fern from our API Definition.

package issues

import (
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	filestorage "github.com/merge-api/merge-go-client/filestorage"
	http "net/http"
	url "net/url"
	time "time"
)

type Client interface {
	List(ctx context.Context, request *filestorage.IssuesListRequest) (*filestorage.PaginatedIssueList, error)
	Retrieve(ctx context.Context, id string) (*filestorage.Issue, error)
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

// Gets issues.
func (c *client) List(ctx context.Context, request *filestorage.IssuesListRequest) (*filestorage.PaginatedIssueList, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/filestorage/v1/issues"

	queryParams := make(url.Values)
	if request.AccountToken != nil {
		queryParams.Add("account_token", fmt.Sprintf("%v", *request.AccountToken))
	}
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.EndDate != nil {
		queryParams.Add("end_date", fmt.Sprintf("%v", *request.EndDate))
	}
	if request.EndUserOrganizationName != nil {
		queryParams.Add("end_user_organization_name", fmt.Sprintf("%v", *request.EndUserOrganizationName))
	}
	if request.FirstIncidentTimeAfter != nil {
		queryParams.Add("first_incident_time_after", fmt.Sprintf("%v", request.FirstIncidentTimeAfter.Format(time.RFC3339)))
	}
	if request.FirstIncidentTimeBefore != nil {
		queryParams.Add("first_incident_time_before", fmt.Sprintf("%v", request.FirstIncidentTimeBefore.Format(time.RFC3339)))
	}
	if request.IncludeMuted != nil {
		queryParams.Add("include_muted", fmt.Sprintf("%v", *request.IncludeMuted))
	}
	if request.IntegrationName != nil {
		queryParams.Add("integration_name", fmt.Sprintf("%v", *request.IntegrationName))
	}
	if request.LastIncidentTimeAfter != nil {
		queryParams.Add("last_incident_time_after", fmt.Sprintf("%v", request.LastIncidentTimeAfter.Format(time.RFC3339)))
	}
	if request.LastIncidentTimeBefore != nil {
		queryParams.Add("last_incident_time_before", fmt.Sprintf("%v", request.LastIncidentTimeBefore.Format(time.RFC3339)))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.StartDate != nil {
		queryParams.Add("start_date", fmt.Sprintf("%v", *request.StartDate))
	}
	if request.Status != nil {
		queryParams.Add("status", fmt.Sprintf("%v", *request.Status))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *filestorage.PaginatedIssueList
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
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

// Get a specific issue.
func (c *client) Retrieve(ctx context.Context, id string) (*filestorage.Issue, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/filestorage/v1/issues/%v", id)

	var response *filestorage.Issue
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
