package testing

import (
	"fmt"
	"github.com/huaweicloud/golangsdk/openstack/deh/v1/hosts"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/openstack/deh/v1/common"
	"testing"
	"net/http"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)


func TestAllocateDeH(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/dedicated-hosts", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, allocateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, allocateResponse)
	})

	c := client.ServiceClient()
	allocateOpts := hosts.AllocateOpts{Name: "Test-1",
		AvailabilityZone: "eu-de-02",
		HostType:         "h1",
		AutoPlacement:    "off",
		Quantity:         2}
	s, err := hosts.Allocate(c, allocateOpts).Extract()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, &hosts.AllocatedHosts{
		AllocatedHostIds: []string{"fb4733fd-70a3-44e1-a1cb-0311f028d7e5",
			"7408f985-047d-4313-b3c8-8e12bef01d12"},
	})
}

func TestUpdateDeH(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/dedicated-hosts/"+HostID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, updateRequest)
		w.WriteHeader(http.StatusNoContent)
	})

	c := client.ServiceClient()
	updateOpts := hosts.UpdateOpts{Name: "Test-2",
		AutoPlacement: "off",
	}
	s := hosts.Update(c, HostID, updateOpts)
	th.AssertNoErr(t, s.Err)
}

func TestDeleteDeH(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/dedicated-hosts/"+HostID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})

	result := hosts.Delete(client.ServiceClient(), HostID)
	th.AssertNoErr(t, result.Err)
}
