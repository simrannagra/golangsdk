package v1

import (
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/rts/v1/stackresources"
	"github.com/huaweicloud/golangsdk/openstack/rts/v1/stacks"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	th "github.com/huaweicloud/golangsdk/testhelper"
	/*"os"
	"fmt"*/
)

func TestStackResources(t *testing.T) {
	// Create a provider client for making the HTTP requests.
	// See common.go in this directory for more information.
	client := newClient(t)

	stackName := "postman_stack_2"
	templateopts:=new(stacks.Template)
	templateopts.Bin=[]byte(template)
	createOpts := stacks.CreateOpts{
		Name:     stackName,
		TemplateOpts: templateopts,
		Timeout:  5,
	}
	stack, err := stacks.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created stack: %+v\n", stack)
	defer func() {
		err := stacks.Delete(client, stackName, stack.ID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted stack (%s)", stackName)
	}()
	err = golangsdk.WaitFor(60, func() (bool, error) {
		getStack, err := stacks.Get(client, stackName, stack.ID).Extract()
		if err != nil {
			return false, err
		}
		if getStack.Status == "CREATE_COMPLETE" {
			return true, nil
		}
		return false, nil
	})


	allResource,err := stackresources.List(client, stackName,stack.ID, stackresources.ListOpts{})
	if err != nil {
		t.Fatalf("Unable to list vpcs: %v", err)
	}
	for _, resource := range allResource {
		tools.PrintResource(t, resource)
	}

}
