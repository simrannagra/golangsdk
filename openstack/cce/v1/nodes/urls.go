package nodes

import "github.com/huaweicloud/golangsdk"

const (
	apiName      = "clusters"
	resourcePath = "hosts"
)

func rootURL(c *golangsdk.ServiceClient, clusteruuid string) string {
	return c.ServiceURL(apiName, clusteruuid, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, clusteruuid string, hostuuid string) string {
	return c.ServiceURL(apiName, clusteruuid, resourcePath, hostuuid)
}

