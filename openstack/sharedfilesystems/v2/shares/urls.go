package shares

import "github.com/huaweicloud/golangsdk"

//For access rule create , update and delete
func rootURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}

// To fetch mount locations of the specified share id
func getMountLocationsURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "export_locations")
}

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("shares")
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id)
}


