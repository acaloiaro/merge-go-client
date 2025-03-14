// This file was auto-generated by Fern from our API Definition.

package hris

import (
	fmt "fmt"
	time "time"
)

type LocationsListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-"`
	// If provided, will only return locations with this location_type
	//
	// - `HOME` - HOME
	// - `WORK` - WORK
	LocationType *LocationsListRequestLocationType `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *LocationsListRequestRemoteFields `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
	ShowEnumOrigins *LocationsListRequestShowEnumOrigins `json:"-"`
}

type LocationsRetrieveRequest struct {
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *LocationsRetrieveRequestRemoteFields `json:"-"`
	// A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
	ShowEnumOrigins *LocationsRetrieveRequestShowEnumOrigins `json:"-"`
}

type LocationsListRequestLocationType string

const (
	LocationsListRequestLocationTypeHome LocationsListRequestLocationType = "HOME"
	LocationsListRequestLocationTypeWork LocationsListRequestLocationType = "WORK"
)

func NewLocationsListRequestLocationTypeFromString(s string) (LocationsListRequestLocationType, error) {
	switch s {
	case "HOME":
		return LocationsListRequestLocationTypeHome, nil
	case "WORK":
		return LocationsListRequestLocationTypeWork, nil
	}
	var t LocationsListRequestLocationType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LocationsListRequestLocationType) Ptr() *LocationsListRequestLocationType {
	return &l
}

type LocationsListRequestRemoteFields string

const (
	LocationsListRequestRemoteFieldsCountry             LocationsListRequestRemoteFields = "country"
	LocationsListRequestRemoteFieldsCountryLocationType LocationsListRequestRemoteFields = "country,location_type"
	LocationsListRequestRemoteFieldsLocationType        LocationsListRequestRemoteFields = "location_type"
)

func NewLocationsListRequestRemoteFieldsFromString(s string) (LocationsListRequestRemoteFields, error) {
	switch s {
	case "country":
		return LocationsListRequestRemoteFieldsCountry, nil
	case "country,location_type":
		return LocationsListRequestRemoteFieldsCountryLocationType, nil
	case "location_type":
		return LocationsListRequestRemoteFieldsLocationType, nil
	}
	var t LocationsListRequestRemoteFields
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LocationsListRequestRemoteFields) Ptr() *LocationsListRequestRemoteFields {
	return &l
}

type LocationsListRequestShowEnumOrigins string

const (
	LocationsListRequestShowEnumOriginsCountry             LocationsListRequestShowEnumOrigins = "country"
	LocationsListRequestShowEnumOriginsCountryLocationType LocationsListRequestShowEnumOrigins = "country,location_type"
	LocationsListRequestShowEnumOriginsLocationType        LocationsListRequestShowEnumOrigins = "location_type"
)

func NewLocationsListRequestShowEnumOriginsFromString(s string) (LocationsListRequestShowEnumOrigins, error) {
	switch s {
	case "country":
		return LocationsListRequestShowEnumOriginsCountry, nil
	case "country,location_type":
		return LocationsListRequestShowEnumOriginsCountryLocationType, nil
	case "location_type":
		return LocationsListRequestShowEnumOriginsLocationType, nil
	}
	var t LocationsListRequestShowEnumOrigins
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LocationsListRequestShowEnumOrigins) Ptr() *LocationsListRequestShowEnumOrigins {
	return &l
}

type LocationsRetrieveRequestRemoteFields string

const (
	LocationsRetrieveRequestRemoteFieldsCountry             LocationsRetrieveRequestRemoteFields = "country"
	LocationsRetrieveRequestRemoteFieldsCountryLocationType LocationsRetrieveRequestRemoteFields = "country,location_type"
	LocationsRetrieveRequestRemoteFieldsLocationType        LocationsRetrieveRequestRemoteFields = "location_type"
)

func NewLocationsRetrieveRequestRemoteFieldsFromString(s string) (LocationsRetrieveRequestRemoteFields, error) {
	switch s {
	case "country":
		return LocationsRetrieveRequestRemoteFieldsCountry, nil
	case "country,location_type":
		return LocationsRetrieveRequestRemoteFieldsCountryLocationType, nil
	case "location_type":
		return LocationsRetrieveRequestRemoteFieldsLocationType, nil
	}
	var t LocationsRetrieveRequestRemoteFields
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LocationsRetrieveRequestRemoteFields) Ptr() *LocationsRetrieveRequestRemoteFields {
	return &l
}

type LocationsRetrieveRequestShowEnumOrigins string

const (
	LocationsRetrieveRequestShowEnumOriginsCountry             LocationsRetrieveRequestShowEnumOrigins = "country"
	LocationsRetrieveRequestShowEnumOriginsCountryLocationType LocationsRetrieveRequestShowEnumOrigins = "country,location_type"
	LocationsRetrieveRequestShowEnumOriginsLocationType        LocationsRetrieveRequestShowEnumOrigins = "location_type"
)

func NewLocationsRetrieveRequestShowEnumOriginsFromString(s string) (LocationsRetrieveRequestShowEnumOrigins, error) {
	switch s {
	case "country":
		return LocationsRetrieveRequestShowEnumOriginsCountry, nil
	case "country,location_type":
		return LocationsRetrieveRequestShowEnumOriginsCountryLocationType, nil
	case "location_type":
		return LocationsRetrieveRequestShowEnumOriginsLocationType, nil
	}
	var t LocationsRetrieveRequestShowEnumOrigins
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LocationsRetrieveRequestShowEnumOrigins) Ptr() *LocationsRetrieveRequestShowEnumOrigins {
	return &l
}
