package cluster

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/cce/v1/cluster"
)

func TestClusterList(t *testing.T) {
	client, err := clients.NewCCEV1Client()
	if err != nil {
		t.Fatalf("Unable to create a vpc client: %v", err)
	}

	listOpts := cluster.ListOpts{}
	allClusters, err := cluster.List(client).ExtractCluster(listOpts)
	if err != nil {
		t.Fatalf("Unable to list clusters: %v", err)
	}
	for _, cluster := range allClusters {
		tools.PrintResource(t, cluster)
	}
}
