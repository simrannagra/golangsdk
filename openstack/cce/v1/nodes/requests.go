package nodes

import (
	"github.com/huaweicloud/golangsdk"
	"reflect"
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
	_, r.Err = c.Post(rootURL(c, clusteruuid), b, nil, reqOpt)
	return
}

// Get retrieves a particular node based on its unique ID.
func Get(c *golangsdk.ServiceClient, clusteruuid, hostuuid string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, clusteruuid, hostuuid), &r.Body, &golangsdk.RequestOpts{
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
	_, r.Err = c.Delete(rootURL(c, clusteruuid), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders,
		JSONBody:    opts,
	})
	return
}

// ListOpts allows the filtering  of paginated collections through
// the API.
type ListOpts struct {
	Name 	string     `json:"name"`
	ID 		string     `json:"uuid"`

}

// List returns collection of
// clusters. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those clusters that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(client *golangsdk.ServiceClient, clusteruuid string) (r ListResult) {

	_,r.Err = client.Get(rootURL(client, clusteruuid), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})

	return
}
func FilterNodes(nodes []Hostlist, opts ListOpts) ([]Hostlist, error) {

	var refinedNodes []Hostlist
	var  clusfield string
	var matched bool
	m := map[string]interface{}{}

	if opts.Name != "" {
		clusfield="Metadata"
		m["Name"] = opts.Name
	}
	if opts.ID != "" {
		clusfield="Metadata"
		m["ID"] = opts.ID
	}

	if len(m) > 0 && len(nodes) > 0 {
		for _, node := range nodes {
			matched = true

			for key, value := range m {

				if sVal := getStructNestedField(&node, clusfield, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedNodes = append(refinedNodes, node)
			}

		}

	} else {
		refinedNodes = nodes
	}

	return refinedNodes, nil
}

func getStructNestedField(v *Hostlist, clusfield string , field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(clusfield).Interface()
	r1 := reflect.ValueOf(f)
	f1 := reflect.Indirect(r1).FieldByName(field)
	return string(f1.String())
}

