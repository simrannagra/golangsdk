package hosts

import "github.com/huaweicloud/golangsdk"

// AllocateOptsBuilder allows extensions to add additional parameters to the
// Allocate request.
type AllocateOptsBuilder interface {
	ToDeHAllocateMap() (map[string]interface{}, error)
}

// AllocateOpts contains all the values needed to allocate a new DeH.
type AllocateOpts struct {
	Name             string `json:"name"`
	AvailabilityZone string `json:"availability_zone"`
	AutoPlacement    string `json:"auto_placement"`
	HostType         string `json:"host_type"`
	Quantity         int    `json:"quantity"`
}

// ToDeHAllocateMap builds a allocate request body from AllocateOpts.
func (opts AllocateOpts) ToDeHAllocateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Allocate accepts a AllocateOpts struct and uses the values to allocate a new DeH.
func Allocate(c *golangsdk.ServiceClient, opts AllocateOptsBuilder) (r AllocateResult) {
	b, err := opts.ToDeHAllocateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200, 201}}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, reqOpt)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToDeHUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update a DeH.
type UpdateOpts struct {
	Name          string `json:"name"`
	AutoPlacement string `json:"auto_placement,omitempty"`
}

// ToDeHUpdateMap builds a update request body from UpdateOpts.
func (opts UpdateOpts) ToDeHUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "dedicated_host")
}

// Update accepts a UpdateOpts struct and uses the values to update a DeH.
func Update(c *golangsdk.ServiceClient, hostID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToDeHUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{204}}
	_, r.Err = c.Put(CommonURL(c, hostID), b, nil, reqOpt)
	return
}

//Deletes the DeH using the specified hostID.
func Delete(c *golangsdk.ServiceClient, hostid string) (r DeleteResult) {
	_, r.Err = c.Delete(CommonURL(c, hostid), nil)
	return
}
