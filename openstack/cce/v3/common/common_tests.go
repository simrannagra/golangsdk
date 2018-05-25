package common

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "api/" + "v3/"+"projects/"+"95fb48c35a6c49bca3f99924885d0362/"
	return sc
}
