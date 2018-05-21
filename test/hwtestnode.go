package main
import (

	"github.com/huaweicloud/golangsdk/openstack"
	"github.com/huaweicloud/golangsdk"
	"fmt"
	"github.com/huaweicloud/golangsdk/openstack/cce/v3/nodes"
	//"github.com/huaweicloud/golangsdk/openstack/cce/v1/clusters"
)
func main() {

	opts1 := golangsdk.AuthOptions{

		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		Username:         "Longyin_kk",
		Password:         "kuaiLE02",
		DomainName:       "longyin_kk",
		TenantID:         "95fb48c35a6c49bca3f99924885d0362",
		//TenantName:       "eu-de_Nordea",
		//TokenID:		  "MIIDmQYJKoZIhvcNAQcCoIIDijCCA4YCAQExDTALBglghkgBZQMEAgEwggFnBgkqhkiG9w0BBwGgggFYBIIBVHsidG9rZW4iOnsiaXNzdWVkX2F0IjoiMjAxNy0xMS0wNlQwNDoyMzo0Ni45ODQwMDBaIiwiZXhwaXJlc19hdCI6IjIwMTctMTEtMDdUMDQ6MjM6NDYuOTg0MDAwWiIsIm1ldGhvZHMiOlsicGFzc3dvcmQiXSwidXNlciI6eyJkb21haW4iOnsibmFtZSI6Ik9UQzAwMDAwMDAwMDAxMDAwMDEwNTA3IiwiaWQiOiI3MTYyOTA3MzgxOGI0NDVkOGQzMDU5MWM1NDUwNTZkMCIsInhkb21haW5fdHlwZSI6IlRTSSIsInhkb21haW5faWQiOiIwMDAwMDAwMDAwMTAwMDAxMDUwNyJ9LCJpZCI6IjczY2YxOTVkMjFjNzQzNjM5NzVkMTllMmIwMGNjMDk1IiwibmFtZSI6ImxpeW9uZ2xlIn0sImNhdGFsb2ciOltdfX0xggIFMIICAQIBATBcMFcxCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVVbnNldDEOMAwGA1UEBwwFVW5zZXQxDjAMBgNVBAoMBVVuc2V0MRgwFgYDVQQDDA93d3cuZXhhbXBsZS5jb20CAQEwCwYJYIZIAWUDBAIBMA0GCSqGSIb3DQEBAQUABIIBgJeSqcO-SQuw1zyoxkFMFwG71JYrQk439VMKSgC+CbGR3uYN6AauNimQKSgE42zWA+t0uGBXea4K5wS183iEumCElBDnr4qtsTv0zNBPA81lfzKxzRzvIwxp9qrOJO4xKY4B4avo4UXHuyCTxqaIr8qc1gfnBudWnCQkRjCJccQazs9iD4NaNUkwS2GiVLm2vQs-xVdjcSoNx1WRMdL2pHk-VpsM-NLC8WPioVOZrTJlbt53tUb5EOFwysxbsAYiaIjthqWuAlziRyPE5by-hibchucmoggwec1T99IZjAaBfMu0XkKobDnfGiHv1np+auTABKwQK+J8+B2PsqtqbQW1MQxUTeooF4YuK-6Ekv6Q14ybogJEhq6yQOwV99Gbv+aGlxDwRt1zFYwzFpLC824qeAdp68fvHmL9Zlu0Ja0bsViuUvQzoRmquETmaAYu+9O9aWCAJpAR8169BxKh5GU-7O5yz0l4t9eFPcy4imCDqyHDGV7pmuNaIOC6MxRqww==",
	}

	provider, err := openstack.AuthenticatedClient(opts1)
	if err != nil {
		fmt.Println(err)
	}

	endpoint := golangsdk.EndpointOpts{
		Region: "cn-east-2",
	}
	client, err := openstack.NewCCEHWV1(provider, endpoint)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(client)


	createNode:=nodes.CreateOpts{
		Kind:"Node",
		Apiversion:"v3" ,
		Metadata:nodes.MetadataOpts{Name:"test-node-1"},
		Spec:nodes.SpecOpts{Flavor: "s1.medium", Az: "cn-east-2a",
							Login:nodes.LoginOpts{SshKey: "c2c-keypair"},
							RootVolume: nodes.VolumeOpts{Size: 40, Volumetype: "SATA"},
							DataVolumes:[]nodes.VolumeDataOpts{{Size: 100, Volumetype:"SATA"}, {Size: 100, Volumetype: "SATA"}},
							//DataVolumes:nodes.VolumeDataOpts{Size: 100, Volumetype: "SATA"},
							Count: 1},
							}
	out,err:=nodes.Create(client,"cec124c2-58f1-11e8-ad73-0255ac101926",createNode).Extract()
	fmt.Println(out)
	fmt.Println(err)

	/*result:=nodes.Get(client,"cec124c2-58f1-11e8-ad73-0255ac101926", "cf4bc001-58f1-11e8-ad73-0255ac101926")
	outget,err:=result.Extract()
	fmt.Println(outget)*/

	/*listnode:=nodes.ListOpts{}
	out,err:=nodes.List(client,"cec124c2-58f1-11e8-ad73-0255ac101926").ExtractNode(listnode)
	fmt.Println(out[0])*/

	/*out:=nodes.Delete(client,"cec124c2-58f1-11e8-ad73-0255ac101926", "4de574eb-599a-11e8-89f3-0255ac101c28")
	fmt.Println(out)*/

	/*updatenode:=nodes.UpdateOpts{Metadata:nodes.UpdateMetadata{Name: "test-node-up"}}
	out,err:=nodes.Update(client,"cec124c2-58f1-11e8-ad73-0255ac101926","d15c02cd-5990-11e8-ad73-0255ac101926",updatenode).Extract()
	fmt.Println(out)*/

}
