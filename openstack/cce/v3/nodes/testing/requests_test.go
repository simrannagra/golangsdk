package testing

import (
	"fmt"
	fake "github.com/huaweicloud/golangsdk/openstack/cce/v3/common"
	"github.com/huaweicloud/golangsdk/openstack/cce/v3/nodes"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"net/http"
	"testing"
)

func TestGetV3Node(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/95fb48c35a6c49bca3f99924885d0362/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/nodes/cf4bc001-58f1-11e8-ad73-0255ac101926", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, Output)
	})

	actual, err := nodes.Get(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", "cf4bc001-58f1-11e8-ad73-0255ac101926").Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)

}

func TestCreateV3Node(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/95fb48c35a6c49bca3f99924885d0362/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/nodes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `
			{
	  "apiversion": "v3",
	  "kind": "Node",
	  "metadata": {
	    "name": "c2c-hostname"
	  },
	  "spec": {
	    "az": "cn-east-2a",
	    "count": 1,
	    "dataVolumes": [
	      {
	        "size": 100,
	        "volumetype": "SATA"
	      }
	    ],
	    "flavor": "s3.large.2",
	    "login": {
	      "sshKey": "c2c-keypair"
	    },
	    "publicIP": {
	      "eip": {
	        "bandwidth": {
	          "sharetype": "",
	          "size": 0
	        },
	        "iptype": ""
	      }
	    },
	    "rootVolume": {
	      "size": 40,
	      "volumetype": "SATA"
	    }
	  }
	}

`)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, Output)
	})
	options := nodes.CreateOpts{Kind: "Node",
		ApiVersion: "v3",
		Metadata:   nodes.CreateMetaData{Name: "c2c-hostname"},
		Spec: nodes.Spec{Flavor: "s3.large.2", Az: "cn-east-2a",
			Login:       nodes.LoginSpec{SshKey: "c2c-keypair"},
			RootVolume:  nodes.VolumeSpec{Size: 40, VolumeType: "SATA"},
			DataVolumes: []nodes.VolumeSpec{{Size: 100, VolumeType: "SATA"}},
			Count:       1},
	}
	actual, err := nodes.Create(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", options).Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)

}

func TestUpdateV3Node(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/95fb48c35a6c49bca3f99924885d0362/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/nodes/cf4bc001-58f1-11e8-ad73-0255ac101926", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "metadata": {
        "name": "c2c-hostname"
    }
}			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, Output)
	})
	options := nodes.UpdateOpts{Metadata: nodes.UpdateMetadata{Name: "c2c-hostname"}}
	actual, err := nodes.Update(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", "cf4bc001-58f1-11e8-ad73-0255ac101926", options).Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)
}

func TestDeleteNode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/95fb48c35a6c49bca3f99924885d0362/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/nodes/cf4bc001-58f1-11e8-ad73-0255ac101926", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
	})

	err := nodes.Delete(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", "cf4bc001-58f1-11e8-ad73-0255ac101926").ExtractErr()
	th.AssertNoErr(t, err)

}
