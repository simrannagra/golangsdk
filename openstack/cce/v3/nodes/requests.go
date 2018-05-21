package nodes

import (
	"github.com/huaweicloud/golangsdk"
	//"github.com/huaweicloud/golangsdk/pagination"
	//"reflect"
)

var RequestOpts golangsdk.RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json"},
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.

type CreateOpts struct {
	Kind       string       `json:"kind" required:"true"`
	Apiversion string       `json:"apiversion" required:"true"`
	Metadata   MetadataOpts `json:"metadata" required:"true"`
	Spec       SpecOpts     `json:"spec" required:"true"`
	//Status     ListStatus    `json:"status"`
}
type MetadataOpts struct {
	Name        string              `json:"name" required:"true"`
	Labels      []map[string]string `json:"labels,omitempty"`
	Annotations []map[string]string `json:"annotations,omitempty"`
}
type SpecOpts struct {
	Flavor      string           `json:"flavor" required:"true"`
	Az          string           `json:"az" required:"true"`
	Login       LoginOpts        `json:"login" required:"true"`
	RootVolume  VolumeOpts       `json:"rootVolume,omitempty" required:"true"`
	DataVolumes []VolumeDataOpts `json:"dataVolumes" required:"true"`
	PublicIP    PublicIPOpts     `json:"publicIP"`
	BillingMode int              `json:"billingMode"`
	Count       int              `json:"count" required:"true"`
	ExtendParam []string         `json:"extendParam"`
}
type LoginOpts struct {
	SshKey string `json:"sshKey,omitempty" required:"true"`
}
type VolumeOpts struct {
	Size        int      `json:"size" required:"true"`
	Volumetype  string   `json:"volumetype" required:"true"`
	ExtendParam []string `json:"extendParam,omitempty"`
}
type VolumeDataOpts struct {
	Size        int    `json:"size" required:"true"`
	Volumetype  string `json:"volumetype" required:"true"`
	ExtendParam string `json:"extendParam,omitempty"`
}
type PublicIPOpts struct {
	Ids   string  `json:"ids"`
	Count int     `json:"count"`
	Eip   EipOpts `json:"eip"`
}
type EipOpts struct {
	IpType    string        `json:"iptype" required:"true"`
	Bandwidth BandwidthOpts `json:"bandwidth" required:"true"`
}
type BandwidthOpts struct {
	ChargeMode string `json:"chargemode"`
	Size       int    `json:"size" required:"true"`
	ShareType  string `json:"sharetype" required:"true"`
}

func (opts CreateOpts) ToNodeCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// logical routes. When it is created, the routes does not have an internal
// interface - it is not associated to any routes.
//
type CreateOptsBuilder interface {
	ToNodeCreateMap() (map[string]interface{}, error)
}

func Create(c *golangsdk.ServiceClient, clusterid string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToNodeCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{201}}
	_, r.Err = c.Post(rootURL(c, clusterid), b, &r.Body, reqOpt)
	return
}

// Get retrieves a particular vpc based on its unique ID.
func Get(c *golangsdk.ServiceClient, clusterid, nodeid string) (r GetResult) {
	//_, r.Err = c.Get(resourceURL(c, clusterid, nodeid), &r.Body, nil)
	_, r.Err = c.Get(resourceURL(c, clusterid, nodeid), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

type ListOpts struct {
	Kind       string     `json:"kind"`
	Apiversion string     `json:"apiversion"`
	Items      CreateOpts `json:"items"`
}
type ListStatus struct {
	Phase     string `json:"phase"`
	ServerId  string `json:"serverId"`
	PublicIP  string `json:"publicIP"`
	PrivateIP string `json:"privateIP"`
	JobID     string `json:"jobID"`
	Reason    string `json:"reason"`
	Message   string `json:"message"`
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToNodeUpdateMap() (map[string]interface{}, error)
}
type UpdateOpts struct {
	Metadata UpdateMetadata `json:"metadata,omitempty"`
}
type UpdateMetadata struct {
	Name string `json:"name,omitempty"`
}

// ToNodeUpdateMap builds an update body based on UpdateOpts.
func (opts UpdateOpts) ToNodeUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update allows vpcs to be updated. You can update the name, administrative
// state, and the external gateway. For more information about how to set the
// external gateway for a vpc, see Create. This operation does not enable
// the update of vpc interfaces. To do this, use the AddInterface and
// RemoveInterface functions.
func Update(c *golangsdk.ServiceClient, clusterid, nodeid string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToNodeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, clusterid, nodeid), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Delete(c *golangsdk.ServiceClient, clusterid, nodeid string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, clusterid, nodeid), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
func List(client *golangsdk.ServiceClient, clusteruuid string) (r ListResult) {

	_, r.Err = client.Get(rootURL(client, clusteruuid), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})

	return
}
