package nodes

import (
	"github.com/huaweicloud/golangsdk"
)

var RequestOpts golangsdk.RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json"},
}

type CreateOptsBuilder interface {
	ToNodeCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new node.
type CreateOpts struct {
	Kind       string `json:"kind" required:"true"`
	ApiVersion string `json:"apiVersion" required:"true"`
	Spec       Spec   `json:"spec" required:"true"`
	Replicas   int    `json:"replicas" required:"true"`
}

// ToNodeCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToNodeCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// . When it is created, the node does not have an internal
// interface
func Create(c *golangsdk.ServiceClient, clusteruuid string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToNodeCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{201}}
	_, r.Err = c.Post(noderCreateURL(c, clusteruuid), b, nil, reqOpt)
	return
}

// Get retrieves a particular node based on its unique ID.
func Get(c *golangsdk.ServiceClient, clusteruuid, hostuuid string) (r GetResult) {
	_, r.Err = c.Get(noderGetURL(c, clusteruuid, hostuuid), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

type Host struct {
	Name string `json:"name" required:"true"`
}

type RemoveOpts struct {
	Host []Host `json:"hosts" required:"true"`
}

func Delete(c *golangsdk.ServiceClient, clusteruuid string, opts RemoveOpts) (r DeleteResult) {
	_, r.Err = c.Delete(noderCreateURL(c, clusteruuid), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders,
		JSONBody:    opts,
	})
	return
}
