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
	SpaceID string 		`json:"spaceuuid"`
	Flavor string   	`json:"flavor"`
	AZ		string   	`json:"az"`
	SSHKey  string   	`json:"sshkey"`
	Clusteruuid string `json:"clusteruuid"`
	Clustername string `json:"clustername"`
	Privateip   string `json:"privateip"`
	Publicip    string `json:"publicip"`
	Status   string   	`json:"status"`
}

// List returns collection of
// nodes. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those nodes that are owned by the
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
	var matched bool

	m := map[string]FilterMetadata{}

	if opts.Name != "" {
		m["Name"] = FilterMetadata{Value: opts.Name, Driller: []string{"Metadata"},IsDrill:true}
	}
		if opts.ID != "" {
		m["ID"] = FilterMetadata{Value: opts.ID, Driller: []string{"Metadata"},IsDrill:true}
	}

	if opts.Flavor != "" {
		m["Flavor"] = FilterMetadata{Value: opts.Flavor, Driller: []string{"Hostspec"},IsDrill:true}
	}

	if opts.AZ != "" {
		m["AZ"] = FilterMetadata{Value: opts.AZ, Driller: []string{"Hostspec"},IsDrill:true}
	}
	if opts.SSHKey != "" {
		m["SSHKey"] = FilterMetadata{Value: opts.SSHKey, Driller: []string{"Hostspec"},IsDrill:true}
	}
	if opts.Clusteruuid != "" {
		m["Clusteruuid"] = FilterMetadata{Value: opts.Clusteruuid, Driller: []string{"Hostspec"},IsDrill:true}
	}

	if opts.Clustername != "" {
		m["Clustername"] = FilterMetadata{Value: opts.Clustername, Driller: []string{"Hostspec"},IsDrill:true}
	}
	if opts.Privateip != "" {
		m["Privateip"] = FilterMetadata{Value: opts.Privateip, Driller: []string{"Hostspec"},IsDrill:true}
	}
	if opts.Publicip != "" {
		m["Publicip"] = FilterMetadata{Value: opts.Publicip, Driller: []string{"Hostspec"},IsDrill:true}
	}

	if opts.SpaceID != "" {
		m["SpaceID"] = FilterMetadata{Value: opts.SpaceID, Driller: []string{"Hostspec"},IsDrill:true}
	}

	if opts.Status != "" {
		m["Status"] = FilterMetadata{Value: opts.Status, Driller: []string{"Status"},IsDrill:false}
	}

	if len(m) > 0 && len(nodes) > 0 {
		for _, node := range nodes {
			matched = true

			for key, value := range m {

				if sVal := GetStructNestedField(&node, key, value.Driller, value.IsDrill); !(sVal == value.Value) {
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

func GetStructNestedField(v *Hostlist, field string, structDriller []string, IsDrill bool) string {
	r := reflect.ValueOf(v)
	if (IsDrill) {
		for _, drillField := range structDriller {
			f := reflect.Indirect(r).FieldByName(drillField).Interface()
			r = reflect.ValueOf(f)
		}
	}
	f1 := reflect.Indirect(r).FieldByName(field)
	return string(f1.String())
}


type FilterMetadata struct {
	Value   string
	Driller []string
	IsDrill bool
}
