package cluster
import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/cce/v1/cluster"
)

func TestVpcList(t *testing.T) {
	client, err := clients.NewClusterV1Client()
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

/*func TestVpcsCRUD(t *testing.T) {
	client, err := clients.NewClusterV1Client()
	if err != nil {
		t.Fatalf("Unable to create a vpc client: %v", err)
	}

	newClusterCertificate, err := cluster.GetCertificate(client, cluster.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to retrieve cluster: %v", err)
	}

	tools.PrintResource(t, newClusterCertificate)
}
*/