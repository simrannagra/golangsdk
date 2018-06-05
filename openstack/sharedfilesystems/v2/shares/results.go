package shares

import (

	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

// AccessRight contains all information associated with an OpenStack share
// Grant Access Response
type AccessRight struct {
	// The UUID of the share to which you are granted or denied access.
	ShareID string `json:"share_id"`
	// The access rule type that can be "ip", "cert" or "user".
	AccessType string `json:"access_type,omitempty"`
	// The value that defines the access that can be a valid format of IP, cert or user.
	AccessTo string `json:"access_to,omitempty"`
	// The access credential of the entity granted share access.
	AccessKey string `json:"access_key,omitempty"`
	// The access level to the share is either "rw" or "ro".
	AccessLevel string `json:"access_level,omitempty"`
	// The state of the access rule
	State string `json:"state,omitempty"`
	// The access rule ID.
	ID string `json:"id"`
}

type DeleteRight struct {
	// The UUID of the share to which you are granted or denied access.
	AccessID string `json:"access_id"`
}

// Extract will get the GrantAccess object from the commonResult
func (r GrantAccessResult) Extract() (*AccessRight, error) {
	var s struct {
		AccessRight *AccessRight `json:"access"`
	}
	err := r.ExtractInto(&s)
	return s.AccessRight, err
}

// GrantAccessResult contains the result body and error from an GrantAccess request.
type GrantAccessResult struct {
	commonResult
	//golangsdk.Result
}

// Extract will get a slice of AccessRight objects from the commonResult
func (r ListAccessRightsResult) Extract() ([]AccessRight, error) {
	var s struct {
		AccessRights []AccessRight `json:"access_list"`
	}
	err := r.ExtractInto(&s)
	return s.AccessRights, err
}

// ListAccessRightsResult contains the result body and error from a ListAccessRights request.
type ListAccessRightsResult struct {
	golangsdk.Result
}

//DeleteAccessResult contains the response body from DeleteAccess rights
type DeleteAccessResult struct {
	golangsdk.Result
}

//GetExportLocationsResult contains the response body from GetExportLocations
type GetExportLocationsResult struct {
	golangsdk.Result
}

// ExportLocation contains all information associated with a share export location
type ExportLocation struct {
	// The export location path that should be used for mount operation.
	Path string `json:"path"`
	// The UUID of the share instance that this export location belongs to.
	ShareInstanceID string `json:"share_instance_id"`
	// Defines purpose of an export location.
	// If set to true, then it is expected to be used for service needs
	// and by administrators only.
	// If it is set to false, then this export location can be used by end users.
	IsAdminOnly bool `json:"is_admin_only"`
	// The share export location UUID.
	ID string `json:"id"`
	Preferred bool `json:"preferred"`
}

// Extract will get the Export Locations from the commonResult
func (r GetExportLocationsResult) Extract() ([]ExportLocation, error) {
	var s struct {
		ExportLocations []ExportLocation `json:"export_locations"`
	}
	err := r.ExtractInto(&s)
	return s.ExportLocations, err
}

