package clusters

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/cce/v3/clusters"
)

func TestClusterV3List(t *testing.T) {
	client, err := clients.NewCCEV3Client()
	if err != nil {
		t.Fatalf("Unable to create a cce client: %v", err)
	}

	listOpts := clusters.ListOpts{}
	allClusters, err := clusters.List(client).ExtractCluster(listOpts)
	if err != nil {
		t.Fatalf("Unable to list clusters: %v", err)
	}
	for _, cluster := range allClusters {
		tools.PrintResource(t, cluster)
	}
}
