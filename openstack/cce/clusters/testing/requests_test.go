package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/huaweicloud/golangsdk/openstack/cce/common"
	"github.com/huaweicloud/golangsdk/openstack/cce/clusters"
	th "github.com/huaweicloud/golangsdk/testhelper"
)



func TestGetCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v1/clusters/d83af16b-febd-4e52-bfb0-20850072e2cd", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetOutput)
	})

	actual, err := clusters.Get(fake.ServiceClient(), "d83af16b-febd-4e52-bfb0-20850072e2cd").Extract()
	th.AssertNoErr(t, err)
	expected := GetExpected
	th.AssertDeepEquals(t, expected, actual)

}

func TestCreateCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v1/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	})
	options := clusters.CreateOpts{"cluster","v1",clusters.CreateMetadataspec{"test-cluster"},
		clusters.CreateSpec{Vpc: "5232f396-d6cc-4a81-8de3-afd7a7ecdf456",
			Subnet: "b857922b-cff7-42f2-b5ee-02619c0081e5",
			Region: "eu-de",
			ClusterType: "Single"}}
	resp:=clusters.Create(fake.ServiceClient(),options)
	th.AssertNoErr(t, resp.Err)

}

func TestUpdateCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v1/clusters/d83af16b-febd-4e52-bfb0-20850072e2cd", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	options := clusters.UpdateOpts{"cluster","v1",clusters.UpdateSpec{"test","989249a7-87e9-4a65-8643-9ed6adf125ec"}}
	resp:=clusters.Update(fake.ServiceClient(),"d83af16b-febd-4e52-bfb0-20850072e2cd",options)
	th.AssertNoErr(t, resp.Err)
}

func TestDeleteVpc(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v1/clusters/d83af16b-febd-4e52-bfb0-20850072e2cd", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := clusters.Delete(fake.ServiceClient(), "d83af16b-febd-4e52-bfb0-20850072e2cd")
	th.AssertNoErr(t, res.Err)
}
