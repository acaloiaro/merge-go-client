// This file was auto-generated by Fern from our API Definition.

package ticketing

import (
	fmt "fmt"
	time "time"
)

type UsersListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// If provided, will only return users with emails equal to this value (case insensitive).
	EmailAddress *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *UsersListRequestExpand `json:"-"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
}

type UsersRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *UsersRetrieveRequestExpand `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
}

type UsersListRequestExpand string

const (
	UsersListRequestExpandRoles      UsersListRequestExpand = "roles"
	UsersListRequestExpandTeams      UsersListRequestExpand = "teams"
	UsersListRequestExpandTeamsRoles UsersListRequestExpand = "teams,roles"
)

func NewUsersListRequestExpandFromString(s string) (UsersListRequestExpand, error) {
	switch s {
	case "roles":
		return UsersListRequestExpandRoles, nil
	case "teams":
		return UsersListRequestExpandTeams, nil
	case "teams,roles":
		return UsersListRequestExpandTeamsRoles, nil
	}
	var t UsersListRequestExpand
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UsersListRequestExpand) Ptr() *UsersListRequestExpand {
	return &u
}

type UsersRetrieveRequestExpand string

const (
	UsersRetrieveRequestExpandRoles      UsersRetrieveRequestExpand = "roles"
	UsersRetrieveRequestExpandTeams      UsersRetrieveRequestExpand = "teams"
	UsersRetrieveRequestExpandTeamsRoles UsersRetrieveRequestExpand = "teams,roles"
)

func NewUsersRetrieveRequestExpandFromString(s string) (UsersRetrieveRequestExpand, error) {
	switch s {
	case "roles":
		return UsersRetrieveRequestExpandRoles, nil
	case "teams":
		return UsersRetrieveRequestExpandTeams, nil
	case "teams,roles":
		return UsersRetrieveRequestExpandTeamsRoles, nil
	}
	var t UsersRetrieveRequestExpand
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UsersRetrieveRequestExpand) Ptr() *UsersRetrieveRequestExpand {
	return &u
}
