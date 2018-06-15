package hosts

import "github.com/huaweicloud/golangsdk"

const resourcePath = "dedicated-hosts"

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func CommonURL(c *golangsdk.ServiceClient, hostID string) string {
	return c.ServiceURL(resourcePath, hostID)
}
