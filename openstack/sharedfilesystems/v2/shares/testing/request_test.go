package testing

import (
	"fmt"
	"github.com/huaweicloud/golangsdk/openstack/sharedfilesystems/v2/shares"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
	"net/http"
	"testing"
)

func TestListAccessRights(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(shareEndpoint+"/"+shareID+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, listAccessRightsRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, listAccessRightsResponse)
	})

	c := client.ServiceClient()
	// Client c must have Microversion set; minimum supported microversion for Grant Access is 2.7
	c.Microversion = "2.7"

	s, err := shares.ListAccessRights(c, shareID).Extract()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, []shares.AccessRight{
		{
			AccessType:  "cert",
			AccessTo:    "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
			AccessLevel: "rw",
			State:       "active",
			ID:          "5158f095-4c43-49c0-b5a7-c458e85ed8c8",
		},
	})
}

func TestGrantAcessRight(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc(shareEndpoint+"/"+shareID+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, grantAccessRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, grantAccessResponse)
	})

	c := client.ServiceClient()
	// Client c must have Microversion set; minimum supported microversion for Grant Access is 2.7
	c.Microversion = "2.7"

	grantaccOpts := shares.GrantAccessOpts{AccessTo: "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8", AccessType: "cert", AccessLevel: "rw"}
	s, err := shares.GrantAccess(c, shareID, grantaccOpts).Extract()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, &shares.AccessRight{
		ShareID:     "1b8facf8-b822-4349-a033-e078b2a84b7f",
		AccessType:  "cert",
		AccessTo:    "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
		AccessLevel: "rw",
		State:       "new",
		ID:          "fc32500f-fa78-4f06-8caf-06ad7fb9726c",
	})
}

func TestDeleteAcess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc(shareEndpoint+"/"+shareID+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, deleteAccessRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

	})
	c := client.ServiceClient()
	c.Microversion = "2.7"

	res := shares.DeleteAccessOpts{AccessID: "ea07152b-d08b-4f6b-8785-ce64dce52679"}
	s := shares.DeleteAccess(c, shareID, res)

	th.AssertNoErr(t, s.Err)
}

func TestGetExportLocationsSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(shareEndpoint+"/"+shareID+"/export_locations", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getExportLocationsResponse)
	})

	c := client.ServiceClient()
	s, err := shares.GetExportLocations(c, shareID).Extract()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, []shares.ExportLocation{
		{
			Path: "sfs-nas1.eu-de.otc.t-systems.com:/share-d41ee18b",
			ID:   "fab962ba-4b9a-475e-a380-8e856ed3f92d",
		},
	})
}

