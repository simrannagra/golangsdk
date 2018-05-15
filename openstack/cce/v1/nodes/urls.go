package nodes

import "github.com/huaweicloud/golangsdk"

const (
	apiName      = "clusters"
	resourcePath = "hosts"
)

func noderCreateURL(c *golangsdk.ServiceClient, clusteruuid string) string {
	return c.ServiceURL(apiName, clusteruuid, resourcePath)
}

func noderGetURL(c *golangsdk.ServiceClient, clusteruuid string, hostuuid string) string {
	return c.ServiceURL(apiName, clusteruuid, resourcePath, hostuuid)
}
