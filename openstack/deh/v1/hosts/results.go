package hosts

import (
	"github.com/huaweicloud/golangsdk"
)

type commonResult struct {
	golangsdk.Result
}

// AllocateResult represents the result of a allocate operation. Call its Extract
// method to interpret it as a host.
type AllocateResult struct {
	commonResult
}

// Extract is a function that accepts a result and extracts Allocated Hosts.
func (r AllocateResult) Extract() (*AllocatedHosts, error) {
	var response AllocatedHosts
	err := r.ExtractInto(&response)
	return &response, err
}

//AllocatedHosts is the response structure of the allocated DeH
type AllocatedHosts struct {
	AllocatedHostIds []string `json:"dedicated_host_ids"`
}

// AllocateResult represents the result of a allocate operation. Call its Extract
// method to interpret it as a host.
type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	commonResult
}
