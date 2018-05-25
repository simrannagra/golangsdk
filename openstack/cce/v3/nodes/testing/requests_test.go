package testing

import (
	"fmt"
	fake "github.com/huaweicloud/golangsdk/openstack/cce/v3/common"
	"github.com/huaweicloud/golangsdk/openstack/cce/v3/nodes"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"net/http"
	"testing"
)

func TestListNode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/95fb48c35a6c49bca3f99924885d0362/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/nodes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "kind": "List",
    "apiVersion": "v3",
	"items":
	[ 
       {
            "kind": "Host",
            "apiVersion": "v3",
            "metadata": {
                "name": "test-node-1234",
                "uid": "b99acd73-5d7c-11e8-8e76-0255ac101929",
                "creationTimestamp": "2018-05-22 04:58:10.605829241 +0000 UTC",
                "updateTimestamp": "2018-05-22 05:02:26.528912685 +0000 UTC"
            },
            "spec": {
                "flavor": "s1.medium",
                "az": "cn-east-2a",
                "login": {
                    "sshKey": "c2c-keypair",
                    "userPassword": {}
                },
                "rootVolume": {
                    "volumetype": "SATA",
                    "size": 40
                },
                "dataVolumes": [
                    {
                        "volumetype": "SATA",
                        "size": 100
                    }
                ],
                "publicIP": {
                    "eip": {
                        "bandwidth": {}
                    }
                },
                "billingMode": 0
            },
            "status": {
                "phase": "Active",
                "serverId": "41748e56-33d4-46a1-aa57-2c8c29907995",
                "privateIP": "192.168.0.3"
            }
        }
	]
}
		`)
	})

	listNodes := nodes.ListOpts{Name: "test-node-1234"}
	actual, err := nodes.List(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926").ExtractNode(listNodes)

	if err != nil {
		t.Errorf("Failed to extract nodes: %v", err)
	}

	expected := []nodes.Items{
		{
			Kind:       "Host",
			Apiversion: "v3",
			Metadata: nodes.Metadata{Name: "test-node-1234",
				Uid:               "b99acd73-5d7c-11e8-8e76-0255ac101929",
				CreationTimestamp: "2018-05-22 04:58:10.605829241 +0000 UTC",
				UpdateTimestamp:   "2018-05-22 05:02:26.528912685 +0000 UTC"},
			Spec: nodes.Spec{Az: "cn-east-2a",
				Login:       nodes.Login{SshKey: "c2c-keypair"},
				RootVolume:  nodes.Volume{Size: 40, Volumetype: "SATA"},
				BillingMode: 0,
				DataVolumes: []nodes.Volume{
					{
						Volumetype: "SATA",
						Size:       100,
					}},
				Flavor: "s1.medium",
			},
			Status: nodes.Status{Phase: "Active", ServerID: "41748e56-33d4-46a1-aa57-2c8c29907995", PrivateIP: "192.168.0.3"},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}
