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

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToShareCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains the options for create a Share. This object is
// passed to shares.Create(). For more information about these parameters,
// please refer to the Share object, or the shared file systems API v2
// documentation
type CreateOpts struct {
	// Defines the share protocol to use
	ShareProto string `json:"share_proto" required:"true"`
	// Size in GB
	Size int `json:"size" required:"true"`
	// Defines the share name
	Name string `json:"name,omitempty"`
	// Share description
	Description string `json:"description,omitempty"`
	// ShareType defines the sharetype. If omitted, a default share type is used
	ShareType string `json:"share_type,omitempty"`
	// VolumeType is deprecated but supported. Either ShareType or VolumeType can be used
	VolumeType string `json:"volume_type,omitempty"`
	// The UUID from which to create a share
	SnapshotID string `json:"snapshot_id,omitempty"`
	// Determines whether or not the share is public
	IsPublic *bool `json:"is_public,omitempty"`
	// Key value pairs of user defined metadata
	Metadata map[string]string `json:"metadata,omitempty"`
	// The UUID of the share network to which the share belongs to
	ShareNetworkID string `json:"share_network_id,omitempty"`
	// The UUID of the consistency group to which the share belongs to
	ConsistencyGroupID string `json:"consistency_group_id,omitempty"`
	// The availability zone of the share
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

// ToShareCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToShareCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "share")
}

// Create will create a new Share based on the values in CreateOpts. To extract
// the Share object from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToShareCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
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
	// Delete will delete an existing Share with the given UUID.
	func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
	}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToShareUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains the values used when updating a Share.
type UpdateOpts struct {
	// DisplayName is equivalent to Name. The API supports using both
	// This is an inherited attribute from the block storage API
	DisplayName string `json:"display_name" required:"true"`
	// DisplayDescription is equivalent to Description. The API supports using bot
	// This is an inherited attribute from the block storage API
	DisplayDescription string `json:"display_description,omitempty"`
}

// ToSubnetUpdateMap builds an update body based on UpdateOpts.
func (opts UpdateOpts) ToShareUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "share")
}

// Update allows shares to be updated. You can update the DisplayName, DisplayDescription.
func Update(c *golangsdk.ServiceClient,id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToShareUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}


	//Gets the Mount/Export Locations of the SFS specified
	func GetExportLocations(client *golangsdk.ServiceClient, id string) (r GetExportLocationsResult) {
	client.Microversion="2.9"
	_, r.Err = client.Get(getMountLocationsURL(client, id), &r.Body, nil)
	return
	}

// Get will get a single share with given UUID
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)

	return
}
