package nodes

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type Metadata struct {
	Name    string `json:"name"`
	ID      string `json:"uuid"`
	SpaceID string `json:"spaceuuid"`
}
type Node struct {
	Kind       string   `json:"kind"`
	ApiVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Replicas   int      `json:"replicas"`
	Status     string   `json:"status"`
}
type Volume struct {
	DiskType   string `json:"diskType" required:"true"`
	DiskSize   int    `json:"diskSize" required:"true"`
	VolumeType string `json:"volumeType" required:"true"`
}
type Spec struct {
	Flavor string            `json:"flavor" required:"true"`
	Label  string            `json:"label"`
	Volume []Volume          `json:"volume" required:"true"`
	SSHKey string            `json:"sshkey" required:"true"`
	Snat   bool              `json:"snat omitempty"`
	AZ     string            `json:"az"`
	Tags   map[string]string `json:"tags omitempty"`

	Clusteruuid string `json:"clusteruuid"`
	Clustername string `json:"clustername"`
	Privateip   string `json:"privateip"`
	Publicip    string `json:"publicip"`
	Cpu         int    `json:"cpu"`
	Memory      int    `json:"memory"`
	Hostid      string `json:"hostid"`
}

// NodePage is the page returned by a pager when traversing over a
// collection of nodes.
type NodePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of nodes has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r NodePage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:""`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}

func (r GetResult) Extract() (*Node, error) {
	var s Node
	err := r.ExtractInto(&s)
	return &s, err
}

func (r GetResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

type commonResult struct {
	golangsdk.Result
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Node.
type CreateResult struct {
	golangsdk.ErrResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Node.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
