package testing

import (
	"fmt"
	fake "github.com/huaweicloud/golangsdk/openstack/cce/v1/common"
	"github.com/huaweicloud/golangsdk/openstack/cce/v1/nodes"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v1/clusters/38a61610-e91d-4669-b8ae-7fbc52776287/hosts/43ee34fb-49fb-4489-938a-559b1991b256",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, NodeGetResponse)
		})

	sg, err := nodes.Get(fake.ServiceClient(), "38a61610-e91d-4669-b8ae-7fbc52776287", "43ee34fb-49fb-4489-938a-559b1991b256").Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, "host", sg.Kind)
	th.AssertEquals(t, "v1", sg.ApiVersion)
	th.AssertEquals(t, "c2c-test-cluster-node-1", sg.Metadata.Name)
	th.AssertEquals(t, "5e6e7641-c288-40a2-a5af-014cc600d838", sg.Metadata.ID)
	th.AssertEquals(t, "s1.medium", sg.Spec.Flavor)
	th.AssertEquals(t, "eu-de-02", sg.Spec.AZ)
	th.AssertEquals(t, 1, sg.Replicas)
	th.AssertEquals(t, "ACTIVE", sg.Status)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v1/clusters/00ebf544-4db8-4f0c-9359-75115861a63a/hosts", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

	})

	nodevolume := []nodes.Volume{{DiskSize: 40, DiskType: "root", VolumeType: "SAS"}, {DiskSize: 100, DiskType: "data", VolumeType: "SATA"}}
	opts := nodes.CreateOpts{
		Kind:       "host",
		ApiVersion: "v1",
		Spec: nodes.Spec{
			Flavor: "s1.medium",
			Volume: nodevolume,
			SSHKey: "click2cloud-key",
			Snat:   false,
			AZ:     "eu-de-01",
		},
		Replicas: 1,
	}

	createdNode := nodes.Create(fake.ServiceClient(), "00ebf544-4db8-4f0c-9359-75115861a63a", opts).ExtractErr()
	th.AssertNoErr(t, createdNode)
	fmt.Println(createdNode)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v1/clusters/00ebf544-4db8-4f0c-9359-75115861a63a/hosts", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
	})
	name := []nodes.Host{{Name: "c2c-test-cluster-node-3"}}
	rmvOpts := nodes.RemoveOpts{
		Host: name,
	}
	res := nodes.Delete(fake.ServiceClient(), "00ebf544-4db8-4f0c-9359-75115861a63a", rmvOpts)
	th.AssertNoErr(t, res.Err)
}
