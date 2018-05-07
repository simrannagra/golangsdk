package v1

import (
	"testing"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/rts/v1/stacks"
	"github.com/huaweicloud/golangsdk/openstack/rts/v1/stacktemplates"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestStackTemplates(t *testing.T) {
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

	tmpl, err := stacktemplates.Get(client, stackName, stack.ID).Extract()
	th.AssertNoErr(t, err)
	t.Logf("retrieved template: %+v\n", tmpl)

}
