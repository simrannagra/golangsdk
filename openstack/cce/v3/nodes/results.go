package nodes
import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type Node struct {
	Kind       string `json:"kind"`
	Apiversion string `json:"apiversion"`
	Metadata   MetadataOpts   `json:"metadata" required:"true"`
	Spec   		SpecOpts    `json:"spec"`
	Status   		StatusOpts    `json:"status" required:"true"`
}
type StatusOpts struct {
	JobID string `json:"jobID"`
}
// NodePage is the page returned by a pager when traversing over a
// collection of vpcs.
type NodePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of vpcs has reached
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

// IsEmpty checks whether a NodePage struct is empty.
func (r NodePage) IsEmpty() (bool, error) {
	is, err := ExtractNodes(r)
	return len(is) == 0, err
}

// ExtractNodes accepts a Page struct, specifically a NodePage struct,
// and extracts the elements into a slice of Vpc structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractNodes(r pagination.Page) ([]Node, error) {
	var s struct {
		Nodes []Node `json:""`
	}
	err := (r.(NodePage)).ExtractInto(&s)
	return s.Nodes, err
}
// Extract is a function that accepts a result and extracts a vpc.
func (r commonResult) Extract() (*Node, error) {
	var s Node
	err := r.ExtractInto(&s)
	return &s, err
}
/*func (r GetResult) Extract() (*Node, error) {
	var s Node
	err := r.ExtractInto(&s)
	return &s, err
}

func (r GetResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}*/
type commonResult struct {
	golangsdk.Result
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Vpc.
type CreateResult struct {
	commonResult
}
// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Vpc.
type GetResult struct {
	commonResult
}
// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Vpc.
type UpdateResult struct {
	commonResult
}
// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

func (r commonResult) ExtractNode(opts ListOpts) ([]Node, error) {
	var s []Node
	err := r.ExtractInto(&s)
	if err!= nil{
		return nil,err
	}
	return s, err
}


type ListResult struct {
	commonResult
}