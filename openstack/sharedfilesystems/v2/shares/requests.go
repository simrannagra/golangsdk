package shares

import (

	"github.com/huaweicloud/golangsdk"

)

// GrantAccessOptsBuilder allows extensions to add additional parameters to the
// GrantAccess request.
type GrantAccessOptsBuilder interface {
	ToGrantAccessMap() (map[string]interface{}, error)
}

// GrantAccessOpts contains the options for creation of an GrantAccess request.
// For more information about these parameters, please, refer to the shared file systems API v2,
// Share Actions, Grant Access documentation
type GrantAccessOpts struct {
	// The access rule type that can be "ip", "cert" or "user".
	AccessType string `json:"access_type"`
	// The value that defines the access that can be a valid format of IP, cert or user.
	AccessTo string `json:"access_to"`
	// The access level to the share is either "rw" or "ro".
	AccessLevel string `json:"access_level"`
}

// ToGrantAccessMap assembles a request body based on the contents of a
// GrantAccessOpts.
func (opts GrantAccessOpts) ToGrantAccessMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-allow_access")
}

// GrantAccess will grant access to a Share based on the values in GrantAccessOpts. To extract
// the GrantAccess object from the response, call the Extract method on the GrantAccessResult.
// Client must have Microversion set; minimum supported microversion for GrantAccess is 2.7.
func GrantAccess(client *golangsdk.ServiceClient, share_id string, opts GrantAccessOptsBuilder) (r GrantAccessResult) {
	b, err := opts.ToGrantAccessMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client, share_id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// ListAccessRights lists all access rules assigned to a Share based on its id. To extract
// the AccessRight slice from the response, call the Extract method on the ListAccessRightsResult.
// Client must have Microversion set; minimum supported microversion for ListAccessRights is 2.7.
func ListAccessRights(client *golangsdk.ServiceClient, id string) (r ListAccessRightsResult) {
	requestBody := map[string]interface{}{"os-access_list": nil}
	_, r.Err = client.Post(rootURL(client, id), requestBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete the Access Rule
type DeleteAccessOptsBuilder interface {
	ToDeleteAccessMap() (map[string]interface{}, error)
}

type DeleteAccessOpts struct {
	// The access ID to be deleted
	AccessID string `json:"access_id"`
}

func (opts DeleteAccessOpts) ToDeleteAccessMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-deny_access")
}

//Deletes the Access Rule
func DeleteAccess(client *golangsdk.ServiceClient, share_id string, opts DeleteAccessOptsBuilder) (r DeleteAccessResult) {
	b, err := opts.ToDeleteAccessMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client, share_id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

//Gets the Mount/Export Locations of the SFS specified
func GetExportLocations(client *golangsdk.ServiceClient, id string) (r GetExportLocationsResult) {
	client.Microversion="2.9"
	_, r.Err = client.Get(getMountLocationsURL(client, id), &r.Body, nil)
	return
}
