package v1

import (
	"os"
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

var template = (`
 {
   "outputs": {
              "str1": {
                 "description": "The description of the nat server.",
                 "value": {
                   "get_resource": "random"
                 }
           }
            },
            "heat_template_version": "2013-05-23",
            "description": "A HOT template that create a single server and boot from volume.",
            "parameters": {
              "key_name": {
                "type": "string",
                "description": "Name of existing key pair for the instance to be created.",
                "default": "KeyPair-click2cloud"
           }
         },
            "resources": {
               "random": {
                  "type": "OS::Heat::RandomString",
                  "properties": {
                  "length": "6"
                  }
           }
        }
}`)

func newClient(t *testing.T) *golangsdk.ServiceClient {
	ao, err := openstack.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	client, err := openstack.AuthenticatedClient(ao)
	th.AssertNoErr(t, err)

	c, err := openstack.NewOrchestrationV1(client, golangsdk.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	th.AssertNoErr(t, err)
	return c
}
