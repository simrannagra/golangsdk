package testing

import (
	"testing"
	"net/http"
	"fmt"
	fake"github.com/huaweicloud/golangsdk/openstack/cce/v1/common"
	"github.com/huaweicloud/golangsdk/openstack/cce/v1/cluster"
	th"github.com/huaweicloud/golangsdk/testhelper"
)


func TestListCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v1/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
[
{
    "kind": "cluster",
    "apiVersion": "v1",
    "metadata": {
        "name": "terraform-test",
        "uuid": "d83af16b-febd-4e52-bfb0-20850072e2cd",
        "spaceuuid": "80f4acb2-6403-0c6a-1cf2-07d02ea708cc"
    },
	"spec": {
        "az": "eu-de-01",
        "cpu": 1,
        "memory": 4096,
        "vpc": "vpc-terraform",
        "vpcid": "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
        "subnet": "subnet-28a5",
        "endpoint": "https://192.168.0.45:5443",
        "external_endpoint": "https://160.44.197.88:5443",
        "clustertype": "Single",
        "security_group_id": "01042a5b-eef7-46a4-9076-4d35de563f3a",
		"hostList": {
            "kind": "list",
            "apiVersion": "v1",
            "metadata": {},
			 "spec": {
                "hostList": [
                    {
                        "kind": "host",
                        "apiVersion": "v1",
                        "metadata": {
                            "name": "terraform-test-node-1",
                            "uuid": "a535f227-e0fb-402f-94f8-d97354a8254d",
                            "spaceuuid": "80f4acb2-6403-0c6a-1cf2-07d02ea708cc"
                        },
						 "spec": {
                            "clusteruuid": "d83af16b-febd-4e52-bfb0-20850072e2cd",
                            "clustername": "terraform-test",
                            "privateip": "192.168.0.53",
                            "publicip": "160.44.194.122",
                            "flavor": "s1.medium",
                            "cpu": 1,
                            "memory": 4096,
                            "hostid": "a535f227-e0fb-402f-94f8-d97354a8254d",
                            "az": "eu-de-02",
                            "volume": [
                                {
                                    "diskType": "root",
                                    "diskSize": 40,
                                    "volumeType": "SATA"
                                }
                            ],
                            "sshkey": "KeyPair-terraform",
                            "status": {
                                "capacity": {
                                    "cpu": "1",
                                    "memory": "3844320Ki",
                                    "pods": "110"
                                },
                                "allocatable": {
                                    "cpu": "0.452",
                                    "memory": "3317984Ki",
                                    "pods": "110"
                                },
                                "conditions": [
                                    {
                                        "type": "OutOfDisk",
                                        "status": "False",
                                        "lastHeartbeatTime": "2018-05-12T11:09:57Z",
                                        "lastTransitionTime": "2018-05-12T09:40:00Z",
                                        "reason": "KubeletHasSufficientDisk",
                                        "message": "kubelet has sufficient disk space available"
                                    }
                                ],
                                "addresses": [
                                    {
                                        "type": "InternalIP",
                                        "address": "192.168.0.53"
                                    }
                                ],
                                "daemonEndpoints": {
                                    "kubeletEndpoint": {
                                        "Port": 10250
                                    }
                                },
                                "nodeInfo": {
                                    "machineID": "f21bcad2f362498d924ac24efcdd3e2a",
                                    "systemUUID": "69646393-240E-4B63-86DE-1DE5E341F732",
                                    "bootID": "cae34bd3-277f-4a08-b429-6337b1bca8a0",
                                    "kernelVersion": "3.10.0-229.42.1.97.x86_64",
                                    "osImage": "EulerOS V2.0SP1",
                                    "containerRuntimeVersion": "docker://1.11.2",
                                    "kubeletVersion": "cce-1.7.3.17+8857de1a267f8f-dirty",
                                    "kubeProxyVersion": "cce-1.7.3.17+8857de1a267f8f-dirty"
                                },
                                "images": [
                                    {
                                        "names": [
                                            "fluentd:1.11"
                                        ],
                                        "sizeBytes": 617825114
                                    }
                                ]
                            }
                        },
                        "replicas": 1,
                        "status": "ACTIVE",
                        "message": "terraform-test-node-1"
                    }
                ]
            }
			}
    },
    "clusterStatus": {
        "status": "AVAILABLE"
    },
    "k8s_version": "1.7.3"
}	
]`)
	})

	//count := 0

	actual,err := cluster.List(fake.ServiceClient()).ExtractCluster(cluster.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract clusters: %v", err)
	}

	expected := []cluster.RetrievedCluster{
		{
			Kind:         "cluster",
			ApiVersion:   "v1",
			Metadata:     cluster.MetaData{Name:"terraform-test",
			                                   ID:"d83af16b-febd-4e52-bfb0-20850072e2cd",
			                                   SpaceID:"80f4acb2-6403-0c6a-1cf2-07d02ea708cc"},
			Clusterspec:  cluster.ClusterlistSpec{AZ:"eu-de-01",
			                                      CPU:1,
			                                      Memory:4096,
			                                      VPC:"vpc-terraform",
			                                      VpcId:"5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
			                                      Subnet:"subnet-28a5",
			                                      Endpoint:"https://192.168.0.45:5443",
			                                      ExternalEndpoint:"https://160.44.197.88:5443",
			                                      ClusterType:"Single",
			                                      SecurityGroupId:"01042a5b-eef7-46a4-9076-4d35de563f3a",
												  ClusterHostList: cluster.ClusterHostList{Kind:"list",
												                                           ApiVersion:"v1",
												                                           Metadata:cluster.MetaData{},
												                                           HostListSpec:cluster.HostListSpec{HostList: []cluster.HostList{{Kind: "host",
																							   ApiVersion: "v1",
																							   Metadata: cluster.MetaData{Name: "terraform-test-node-1",
																								   ID: "a535f227-e0fb-402f-94f8-d97354a8254d",
																								   SpaceID: "80f4acb2-6403-0c6a-1cf2-07d02ea708cc",
																							   },
																							   Hostspec: cluster.HostSpec{ClusterUuid: "d83af16b-febd-4e52-bfb0-20850072e2cd",
																								   ClusterName: "terraform-test",
																								   PrivateIp: "192.168.0.53",
																								   PublicIp: "160.44.194.122",
																								   Flavor: "s1.medium",
																								   CPU: 1,
																								   Memory: 4096,
																								   HostId: "a535f227-e0fb-402f-94f8-d97354a8254d",
																								   AZ: "eu-de-02",
																								   SshKey: "KeyPair-terraform",
																								   Volume: []cluster.Volume{{DiskType: "root", DiskSize: 40, VolumeType: "SATA"}},
																								   Status: cluster.HostStatus{Capacity: cluster.Capacity{CPU: "1", Memory: "3844320Ki", Pods: "110"},
																									   Allocatable: cluster.Capacity{CPU: "0.452", Memory: "3317984Ki", Pods: "110"},
																									   Conditions: []cluster.Conditions{{Type: "OutOfDisk", Status: "False", Reason: "KubeletHasSufficientDisk"}},
																									   Addresses: []cluster.Addresses{{Type: "InternalIP", Address: "192.168.0.53"}},
																									   DaemonEndpoints: cluster.DaemonEndpoints{KubeletEndpoint: cluster.KubeletEndpoint{Port: 10250}},
																									   NodeInfo: cluster.NodeInfo{MachineID: "f21bcad2f362498d924ac24efcdd3e2a",
																										   SystemUUID: "69646393-240E-4B63-86DE-1DE5E341F732",
																										   BootID: "cae34bd3-277f-4a08-b429-6337b1bca8a0",
																										   KernelVersion: "3.10.0-229.42.1.97.x86_64",
																										   OsImage: "EulerOS V2.0SP1",
																										   ContainerRuntimeVersion: "docker://1.11.2",
																										   KubeletVersion: "cce-1.7.3.17+8857de1a267f8f-dirty",
																										   KubeProxyVersion: "cce-1.7.3.17+8857de1a267f8f-dirty"},
																									   Images: []cluster.Images{{Names: []string{"fluentd:1.11"}, SizeBytes: 617825114}},}},
																							   Replicas: 1,
																							   Status: "ACTIVE",
																							   Message: "terraform-test-node-1",
																							                              }},
																						   },

	                                              },
	},

			ClusterStatus: cluster.OutStatus{Status:"AVAILABLE"},
			K8sVersion:    "1.7.3",
		},

	}

	th.AssertDeepEquals(t, expected, actual)
}


func TestGetCertificate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v1/clusters/ea28a104-9fa0-48e6-8cee-b1a3e4a34957/certificates", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		//th.TestHeader(t, r, "Content-Type", "application/json")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "cluster_uuid": "ea28a104-9fa0-48e6-8cee-b1a3e4a34957",
    "cluster_name": "test_cluster_certificate",
    "endpoint": "https://172.16.0.102:5443",
    "cacrt": "-----BEGIN CERTIFICATE-----\nMIIEHjCCAoagAwIBAgIJAMXGl60Y02HgMA0GCSqGSIb3DQEBCwUAMA8xDTALBgNV\nBAMTBHRlc3QwHhcNMTgwNDI2MDczNjUwWhcNMjgwNDIzMDczNjUwWjAPMQ0wCwYD\nVQQDEwR0ZXN0MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA7G/hFOwl\nwfZNTPKXhPpNvpScb5SJ+yob+MsSTpWeb4hC2oE0ABnXLYu2fBrHJFYmbQeY07pu\nmKk6KaRYunSoWqQHB24jthGfjc0wsg/7rppcCR9LaSzRu9fHtWzaUl/VeOzcKf1P\n4MIaIgTULTRPmyzwnqcOwYMH/qlgLzKctYB+gCkHdRmHSug+k5t3XcY42do3Lk0W\nIFARLeBT2dzEtBiEg9kifD2dN6aXtpqm7+SlKlq3PbhnJiSg84mNiBeM5w6m2VI5\npjcx0dwEi+QtJ6AwgObtdIk8Wzfv9NLvA+8PS/Al4ZG/W/QiRwIMyLAqACybV/BF\nP6FNkX1RQSbwUWEu3k9BEJ8VP17LU2b2OYWFfRBZbNbefQMV5fyMP9CVyecB031z\nZJThbpb/PMv0t3cjLc4947AIty+KlHugdPihHkprgZHDzKfEOySQRW+10cusn3oT\nBcJpYmj1CYCZ2ahiiv3eVuUOFjFgU++EQO/1/UdXRNBfP2g2QdxZPBArAgMBAAGj\nfTB7MB0GA1UdDgQWBBRudkAppRFi10WTEGgaRrnnAx8pQjA/BgNVHSMEODA2gBRu\ndkAppRFi10WTEGgaRrnnAx8pQqETpBEwDzENMAsGA1UEAxMEdGVzdIIJAMXGl60Y\n02HgMAwGA1UdEwQFMAMBAf8wCwYDVR0PBAQDAgEGMA0GCSqGSIb3DQEBCwUAA4IB\ngQCmtkivyCEXvi2JpEco4+Rt7oPHOirqOOG8uobO8y7d5jg+GN1Gevwoq44lXP7w\nIcNTaXQTgX429VgtSxCYeUJ8l1+o6RmExeH0cj1cdtGcQTTZ52UMS1pCmc3YrnOE\nnas4HwJ49AkZn1nmOnLc40Q/0g8d+DlEdrF+7XUJlMOW5pO9KogE+cWShWRbMJrc\ndkmS/QQIYLOJwdn/waZMveV2uVcyIydG1z4q8rmOdNx7iypZay0vt3dDS59QYvk/\n0gvpeYWPB+QIlmR6mB3AUihO5v8yO/YXZIbhfJ83vGB+zRcdBfXzftfN9gYZj6Ca\nb+EDV0vBTZ7ELVksKBmaQW3GKpUY7t8QxxzYKmlYfjkWmTjHEmqebJmlx58imgWE\n4yovjMom3qaO8dtMIudvhTvbWCpo9mMF9jv8548VTuIH51DGGbrsWqzNN1bJVU/z\nGft4XsiSAH9WFnhE7vGej9D3KY0f7ig8SCQiFTEnaXZyUINFl2/7ludHweb1j06K\n6P8=\n-----END CERTIFICATE-----\n",
    "clientcrt": "Certificate:\n    Data:\n        Version: 3 (0x2)\n        Serial Number: 2 (0x2)\n    Signature Algorithm: sha256WithRSAEncryption\n        Issuer: CN=test\n        Validity\n            Not Before: Apr 26 07:36:52 2018 GMT\n            Not After : Apr 23 07:36:52 2028 GMT\n        Subject: CN=kubecfg\n        Subject Public Key Info:\n            Public Key Algorithm: rsaEncryption\n                Public-Key: (3072 bit)\n                Modulus:\n                    00:b9:b5:7f:46:3e:dd:bf:0d:9b:f0:dd:fb:eb:6d:\n                    dd:0e:13:0e:15:b1:7f:81:b3:84:32:7c:74:ee:91:\n                    c2:27:ed:6b:df:a6:66:6b:eb:81:59:59:c0:22:d9:\n                    10:ac:98:8c:78:76:cd:3f:40:a3:f2:92:b0:2c:6b:\n                    af:f0:42:e5:a2:04:a2:fc:a9:b1:34:b5:37:5e:45:\n                    4d:d0:8c:fd:39:8a:70:24:f4:2d:92:91:81:50:b5:\n                    58:ff:ad:24:15:1d:a9:fc:ac:86:51:a7:d2:7b:44:\n                    d2:c8:cc:3c:18:6f:34:ca:b0:94:01:99:d9:54:e0:\n                    e3:4f:7b:f6:34:fa:92:6b:89:0d:ff:43:eb:42:92:\n                    63:24:d6:14:48:1b:5c:f1:c9:61:39:90:0e:8f:d3:\n                    44:11:91:de:3e:e2:82:9d:6e:37:11:3a:71:bb:9a:\n                    8a:bf:fe:a9:ec:82:3d:fe:a3:32:84:85:4d:cd:c5:\n                    0e:33:20:6c:e0:25:b7:98:f0:24:d1:c0:f3:80:6a:\n                    6b:48:4f:a1:b3:5e:ec:e2:a6:23:fa:c6:97:71:21:\n                    4d:54:6f:ed:28:33:96:78:53:80:14:27:62:1b:b5:\n                    ad:62:8c:7f:99:55:1a:65:cf:22:c5:62:db:b3:3e:\n                    ee:fa:9a:cb:bc:31:74:48:2b:d6:00:40:3d:6f:61:\n                    5a:55:0b:22:cb:63:6d:55:88:a6:05:f4:3d:3f:22:\n                    f5:8b:b4:de:49:fe:23:90:a8:f5:2f:7a:dc:e3:d6:\n                    fe:7d:65:f1:d4:4f:81:df:5f:83:53:89:42:70:7e:\n                    fe:bd:20:be:26:f6:ef:02:d1:84:6e:b6:9e:64:78:\n                    4b:5f:6e:17:c8:a9:fd:36:e4:8f:49:40:02:7e:1e:\n                    a9:7c:6b:00:61:e5:20:b9:f0:68:de:85:d5:53:0c:\n                    15:4b:b2:b5:c7:4f:c1:3f:ae:96:f8:fe:e4:fe:f7:\n                    61:42:eb:d9:28:56:84:a0:86:de:70:53:d6:a5:42:\n                    1f:f7:55:ec:3d:d0:93:0d:be:8f\n                Exponent: 65537 (0x10001)\n        X509v3 extensions:\n            X509v3 Basic Constraints: \n                CA:FALSE\n            X509v3 Subject Key Identifier: \n                48:F8:DA:AF:8C:05:9E:D4:8E:23:46:08:F3:BC:4D:4C:72:E6:8C:7D\n            X509v3 Authority Key Identifier: \n                keyid:6E:76:40:29:A5:11:62:D7:45:93:10:68:1A:46:B9:E7:03:1F:29:42\n                DirName:/CN=test\n                serial:C5:C6:97:AD:18:D3:61:E0\n\n            X509v3 Extended Key Usage: \n                TLS Web Client Authentication\n            X509v3 Key Usage: \n                Digital Signature\n    Signature Algorithm: sha256WithRSAEncryption\n         da:50:e4:19:ec:01:cc:12:99:c9:97:3f:ef:91:c6:fa:35:21:\n         f3:55:50:ce:3d:b6:af:5a:dc:1a:d0:79:02:1d:94:c6:21:05:\n         5d:30:62:cd:1c:43:c8:84:a8:cb:66:07:43:b0:d2:db:37:2f:\n         ed:55:e6:87:4d:2a:7e:39:7f:99:95:f2:8b:be:b2:3b:1a:b4:\n         b2:54:5e:93:1e:15:2e:41:9e:a2:a8:ae:e5:13:ba:16:a8:40:\n         cf:75:96:be:77:c6:1a:46:6f:28:c7:78:3b:4a:cb:aa:71:64:\n         52:e1:c6:05:af:1b:06:83:ef:bb:b4:e8:47:c7:69:2a:59:1e:\n         13:07:76:35:a4:df:93:6e:b9:26:8a:87:d1:b9:8a:e1:80:46:\n         1a:1a:a6:45:6d:4c:8a:c0:75:af:78:49:2e:56:ae:6e:04:be:\n         2b:6b:ad:57:e0:14:fb:5e:78:08:90:03:1e:6a:df:ee:80:51:\n         07:97:91:64:15:71:78:e3:8e:12:4e:7e:65:6a:38:ac:3f:5a:\n         c6:db:2e:33:b6:25:84:e5:78:30:97:f5:a4:9f:6a:33:a8:8b:\n         68:f0:c7:91:2b:d5:86:91:e1:00:9d:67:44:1a:5d:8f:0e:4c:\n         d5:9e:ee:57:6e:e3:e3:86:19:3b:77:51:31:8d:fe:81:e0:13:\n         7e:03:20:0e:e8:fb:f4:ca:11:93:87:8f:e1:56:47:61:b4:c7:\n         d0:cf:e4:5e:e6:b7:df:da:9e:73:a8:8d:3a:f1:4c:a7:56:05:\n         db:66:6c:a3:15:dc:c6:17:dc:77:ef:7d:4f:00:3b:8e:0f:1a:\n         bb:ae:07:e4:a7:f7:c2:a4:f5:25:30:38:ea:95:2a:ab:b7:69:\n         07:41:1f:16:ae:5a:4b:67:26:10:80:ce:14:f5:cc:c7:89:80:\n         0e:7b:43:99:8d:74:2f:9e:db:72:29:75:5e:b4:27:05:2f:9b:\n         c4:8e:70:94:88:18:a6:cd:56:da:87:b9:4c:b5:2b:9a:71:61:\n         04:e5:d0:14:0e:34\n-----BEGIN CERTIFICATE-----\nMIIELTCCApWgAwIBAgIBAjANBgkqhkiG9w0BAQsFADAPMQ0wCwYDVQQDEwR0ZXN0\nMB4XDTE4MDQyNjA3MzY1MloXDTI4MDQyMzA3MzY1MlowEjEQMA4GA1UEAxMHa3Vi\nZWNmZzCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBALm1f0Y+3b8Nm/Dd\n++tt3Q4TDhWxf4GzhDJ8dO6Rwifta9+mZmvrgVlZwCLZEKyYjHh2zT9Ao/KSsCxr\nr/BC5aIEovypsTS1N15FTdCM/TmKcCT0LZKRgVC1WP+tJBUdqfyshlGn0ntE0sjM\nPBhvNMqwlAGZ2VTg40979jT6kmuJDf9D60KSYyTWFEgbXPHJYTmQDo/TRBGR3j7i\ngp1uNxE6cbuair/+qeyCPf6jMoSFTc3FDjMgbOAlt5jwJNHA84Bqa0hPobNe7OKm\nI/rGl3EhTVRv7SgzlnhTgBQnYhu1rWKMf5lVGmXPIsVi27M+7vqay7wxdEgr1gBA\nPW9hWlULIstjbVWIpgX0PT8i9Yu03kn+I5Co9S963OPW/n1l8dRPgd9fg1OJQnB+\n/r0gvib27wLRhG62nmR4S19uF8ip/Tbkj0lAAn4eqXxrAGHlILnwaN6F1VMMFUuy\ntcdPwT+ulvj+5P73YULr2ShWhKCG3nBT1qVCH/dV7D3Qkw2+jwIDAQABo4GQMIGN\nMAkGA1UdEwQCMAAwHQYDVR0OBBYEFEj42q+MBZ7UjiNGCPO8TUxy5ox9MD8GA1Ud\nIwQ4MDaAFG52QCmlEWLXRZMQaBpGuecDHylCoROkETAPMQ0wCwYDVQQDEwR0ZXN0\nggkAxcaXrRjTYeAwEwYDVR0lBAwwCgYIKwYBBQUHAwIwCwYDVR0PBAQDAgeAMA0G\nCSqGSIb3DQEBCwUAA4IBgQDaUOQZ7AHMEpnJlz/vkcb6NSHzVVDOPbavWtwa0HkC\nHZTGIQVdMGLNHEPIhKjLZgdDsNLbNy/tVeaHTSp+OX+ZlfKLvrI7GrSyVF6THhUu\nQZ6iqK7lE7oWqEDPdZa+d8YaRm8ox3g7SsuqcWRS4cYFrxsGg++7tOhHx2kqWR4T\nB3Y1pN+TbrkmiofRuYrhgEYaGqZFbUyKwHWveEkuVq5uBL4ra61X4BT7XngIkAMe\nat/ugFEHl5FkFXF4444STn5lajisP1rG2y4ztiWE5Xgwl/Wkn2ozqIto8MeRK9WG\nkeEAnWdEGl2PDkzVnu5XbuPjhhk7d1Exjf6B4BN+AyAO6Pv0yhGTh4/hVkdhtMfQ\nz+Re5rff2p5zqI068UynVgXbZmyjFdzGF9x3731PADuODxq7rgfkp/fCpPUlMDjq\nlSqrt2kHQR8WrlpLZyYQgM4U9czHiYAOe0OZjXQvnttyKXVetCcFL5vEjnCUiBim\nzVbah7lMtSuacWEE5dAUDjQ=\n-----END CERTIFICATE-----\n",
    "clientkey": "-----BEGIN PRIVATE KEY-----\n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\njB8vl/PiudSrzCa2sKK6xZk=\n-----END PRIVATE KEY-----\n"
}
		`)
	})

	n, err := cluster.GetCertificate(fake.ServiceClient(),"ea28a104-9fa0-48e6-8cee-b1a3e4a34957").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "ea28a104-9fa0-48e6-8cee-b1a3e4a34957", n.ClusterID)
	th.AssertEquals(t, "test_cluster_certificate", n.ClusterName)
	th.AssertEquals(t, "https://172.16.0.102:5443", n.EndPoint)
	th.AssertEquals(t, "-----BEGIN CERTIFICATE-----\nMIIEHjCCAoagAwIBAgIJAMXGl60Y02HgMA0GCSqGSIb3DQEBCwUAMA8xDTALBgNV\nBAMTBHRlc3QwHhcNMTgwNDI2MDczNjUwWhcNMjgwNDIzMDczNjUwWjAPMQ0wCwYD\nVQQDEwR0ZXN0MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA7G/hFOwl\nwfZNTPKXhPpNvpScb5SJ+yob+MsSTpWeb4hC2oE0ABnXLYu2fBrHJFYmbQeY07pu\nmKk6KaRYunSoWqQHB24jthGfjc0wsg/7rppcCR9LaSzRu9fHtWzaUl/VeOzcKf1P\n4MIaIgTULTRPmyzwnqcOwYMH/qlgLzKctYB+gCkHdRmHSug+k5t3XcY42do3Lk0W\nIFARLeBT2dzEtBiEg9kifD2dN6aXtpqm7+SlKlq3PbhnJiSg84mNiBeM5w6m2VI5\npjcx0dwEi+QtJ6AwgObtdIk8Wzfv9NLvA+8PS/Al4ZG/W/QiRwIMyLAqACybV/BF\nP6FNkX1RQSbwUWEu3k9BEJ8VP17LU2b2OYWFfRBZbNbefQMV5fyMP9CVyecB031z\nZJThbpb/PMv0t3cjLc4947AIty+KlHugdPihHkprgZHDzKfEOySQRW+10cusn3oT\nBcJpYmj1CYCZ2ahiiv3eVuUOFjFgU++EQO/1/UdXRNBfP2g2QdxZPBArAgMBAAGj\nfTB7MB0GA1UdDgQWBBRudkAppRFi10WTEGgaRrnnAx8pQjA/BgNVHSMEODA2gBRu\ndkAppRFi10WTEGgaRrnnAx8pQqETpBEwDzENMAsGA1UEAxMEdGVzdIIJAMXGl60Y\n02HgMAwGA1UdEwQFMAMBAf8wCwYDVR0PBAQDAgEGMA0GCSqGSIb3DQEBCwUAA4IB\ngQCmtkivyCEXvi2JpEco4+Rt7oPHOirqOOG8uobO8y7d5jg+GN1Gevwoq44lXP7w\nIcNTaXQTgX429VgtSxCYeUJ8l1+o6RmExeH0cj1cdtGcQTTZ52UMS1pCmc3YrnOE\nnas4HwJ49AkZn1nmOnLc40Q/0g8d+DlEdrF+7XUJlMOW5pO9KogE+cWShWRbMJrc\ndkmS/QQIYLOJwdn/waZMveV2uVcyIydG1z4q8rmOdNx7iypZay0vt3dDS59QYvk/\n0gvpeYWPB+QIlmR6mB3AUihO5v8yO/YXZIbhfJ83vGB+zRcdBfXzftfN9gYZj6Ca\nb+EDV0vBTZ7ELVksKBmaQW3GKpUY7t8QxxzYKmlYfjkWmTjHEmqebJmlx58imgWE\n4yovjMom3qaO8dtMIudvhTvbWCpo9mMF9jv8548VTuIH51DGGbrsWqzNN1bJVU/z\nGft4XsiSAH9WFnhE7vGej9D3KY0f7ig8SCQiFTEnaXZyUINFl2/7ludHweb1j06K\n6P8=\n-----END CERTIFICATE-----\n", n.Cacrt)
	th.AssertEquals(t, "Certificate:\n    Data:\n        Version: 3 (0x2)\n        Serial Number: 2 (0x2)\n    Signature Algorithm: sha256WithRSAEncryption\n        Issuer: CN=test\n        Validity\n            Not Before: Apr 26 07:36:52 2018 GMT\n            Not After : Apr 23 07:36:52 2028 GMT\n        Subject: CN=kubecfg\n        Subject Public Key Info:\n            Public Key Algorithm: rsaEncryption\n                Public-Key: (3072 bit)\n                Modulus:\n                    00:b9:b5:7f:46:3e:dd:bf:0d:9b:f0:dd:fb:eb:6d:\n                    dd:0e:13:0e:15:b1:7f:81:b3:84:32:7c:74:ee:91:\n                    c2:27:ed:6b:df:a6:66:6b:eb:81:59:59:c0:22:d9:\n                    10:ac:98:8c:78:76:cd:3f:40:a3:f2:92:b0:2c:6b:\n                    af:f0:42:e5:a2:04:a2:fc:a9:b1:34:b5:37:5e:45:\n                    4d:d0:8c:fd:39:8a:70:24:f4:2d:92:91:81:50:b5:\n                    58:ff:ad:24:15:1d:a9:fc:ac:86:51:a7:d2:7b:44:\n                    d2:c8:cc:3c:18:6f:34:ca:b0:94:01:99:d9:54:e0:\n                    e3:4f:7b:f6:34:fa:92:6b:89:0d:ff:43:eb:42:92:\n                    63:24:d6:14:48:1b:5c:f1:c9:61:39:90:0e:8f:d3:\n                    44:11:91:de:3e:e2:82:9d:6e:37:11:3a:71:bb:9a:\n                    8a:bf:fe:a9:ec:82:3d:fe:a3:32:84:85:4d:cd:c5:\n                    0e:33:20:6c:e0:25:b7:98:f0:24:d1:c0:f3:80:6a:\n                    6b:48:4f:a1:b3:5e:ec:e2:a6:23:fa:c6:97:71:21:\n                    4d:54:6f:ed:28:33:96:78:53:80:14:27:62:1b:b5:\n                    ad:62:8c:7f:99:55:1a:65:cf:22:c5:62:db:b3:3e:\n                    ee:fa:9a:cb:bc:31:74:48:2b:d6:00:40:3d:6f:61:\n                    5a:55:0b:22:cb:63:6d:55:88:a6:05:f4:3d:3f:22:\n                    f5:8b:b4:de:49:fe:23:90:a8:f5:2f:7a:dc:e3:d6:\n                    fe:7d:65:f1:d4:4f:81:df:5f:83:53:89:42:70:7e:\n                    fe:bd:20:be:26:f6:ef:02:d1:84:6e:b6:9e:64:78:\n                    4b:5f:6e:17:c8:a9:fd:36:e4:8f:49:40:02:7e:1e:\n                    a9:7c:6b:00:61:e5:20:b9:f0:68:de:85:d5:53:0c:\n                    15:4b:b2:b5:c7:4f:c1:3f:ae:96:f8:fe:e4:fe:f7:\n                    61:42:eb:d9:28:56:84:a0:86:de:70:53:d6:a5:42:\n                    1f:f7:55:ec:3d:d0:93:0d:be:8f\n                Exponent: 65537 (0x10001)\n        X509v3 extensions:\n            X509v3 Basic Constraints: \n                CA:FALSE\n            X509v3 Subject Key Identifier: \n                48:F8:DA:AF:8C:05:9E:D4:8E:23:46:08:F3:BC:4D:4C:72:E6:8C:7D\n            X509v3 Authority Key Identifier: \n                keyid:6E:76:40:29:A5:11:62:D7:45:93:10:68:1A:46:B9:E7:03:1F:29:42\n                DirName:/CN=test\n                serial:C5:C6:97:AD:18:D3:61:E0\n\n            X509v3 Extended Key Usage: \n                TLS Web Client Authentication\n            X509v3 Key Usage: \n                Digital Signature\n    Signature Algorithm: sha256WithRSAEncryption\n         da:50:e4:19:ec:01:cc:12:99:c9:97:3f:ef:91:c6:fa:35:21:\n         f3:55:50:ce:3d:b6:af:5a:dc:1a:d0:79:02:1d:94:c6:21:05:\n         5d:30:62:cd:1c:43:c8:84:a8:cb:66:07:43:b0:d2:db:37:2f:\n         ed:55:e6:87:4d:2a:7e:39:7f:99:95:f2:8b:be:b2:3b:1a:b4:\n         b2:54:5e:93:1e:15:2e:41:9e:a2:a8:ae:e5:13:ba:16:a8:40:\n         cf:75:96:be:77:c6:1a:46:6f:28:c7:78:3b:4a:cb:aa:71:64:\n         52:e1:c6:05:af:1b:06:83:ef:bb:b4:e8:47:c7:69:2a:59:1e:\n         13:07:76:35:a4:df:93:6e:b9:26:8a:87:d1:b9:8a:e1:80:46:\n         1a:1a:a6:45:6d:4c:8a:c0:75:af:78:49:2e:56:ae:6e:04:be:\n         2b:6b:ad:57:e0:14:fb:5e:78:08:90:03:1e:6a:df:ee:80:51:\n         07:97:91:64:15:71:78:e3:8e:12:4e:7e:65:6a:38:ac:3f:5a:\n         c6:db:2e:33:b6:25:84:e5:78:30:97:f5:a4:9f:6a:33:a8:8b:\n         68:f0:c7:91:2b:d5:86:91:e1:00:9d:67:44:1a:5d:8f:0e:4c:\n         d5:9e:ee:57:6e:e3:e3:86:19:3b:77:51:31:8d:fe:81:e0:13:\n         7e:03:20:0e:e8:fb:f4:ca:11:93:87:8f:e1:56:47:61:b4:c7:\n         d0:cf:e4:5e:e6:b7:df:da:9e:73:a8:8d:3a:f1:4c:a7:56:05:\n         db:66:6c:a3:15:dc:c6:17:dc:77:ef:7d:4f:00:3b:8e:0f:1a:\n         bb:ae:07:e4:a7:f7:c2:a4:f5:25:30:38:ea:95:2a:ab:b7:69:\n         07:41:1f:16:ae:5a:4b:67:26:10:80:ce:14:f5:cc:c7:89:80:\n         0e:7b:43:99:8d:74:2f:9e:db:72:29:75:5e:b4:27:05:2f:9b:\n         c4:8e:70:94:88:18:a6:cd:56:da:87:b9:4c:b5:2b:9a:71:61:\n         04:e5:d0:14:0e:34\n-----BEGIN CERTIFICATE-----\nMIIELTCCApWgAwIBAgIBAjANBgkqhkiG9w0BAQsFADAPMQ0wCwYDVQQDEwR0ZXN0\nMB4XDTE4MDQyNjA3MzY1MloXDTI4MDQyMzA3MzY1MlowEjEQMA4GA1UEAxMHa3Vi\nZWNmZzCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBALm1f0Y+3b8Nm/Dd\n++tt3Q4TDhWxf4GzhDJ8dO6Rwifta9+mZmvrgVlZwCLZEKyYjHh2zT9Ao/KSsCxr\nr/BC5aIEovypsTS1N15FTdCM/TmKcCT0LZKRgVC1WP+tJBUdqfyshlGn0ntE0sjM\nPBhvNMqwlAGZ2VTg40979jT6kmuJDf9D60KSYyTWFEgbXPHJYTmQDo/TRBGR3j7i\ngp1uNxE6cbuair/+qeyCPf6jMoSFTc3FDjMgbOAlt5jwJNHA84Bqa0hPobNe7OKm\nI/rGl3EhTVRv7SgzlnhTgBQnYhu1rWKMf5lVGmXPIsVi27M+7vqay7wxdEgr1gBA\nPW9hWlULIstjbVWIpgX0PT8i9Yu03kn+I5Co9S963OPW/n1l8dRPgd9fg1OJQnB+\n/r0gvib27wLRhG62nmR4S19uF8ip/Tbkj0lAAn4eqXxrAGHlILnwaN6F1VMMFUuy\ntcdPwT+ulvj+5P73YULr2ShWhKCG3nBT1qVCH/dV7D3Qkw2+jwIDAQABo4GQMIGN\nMAkGA1UdEwQCMAAwHQYDVR0OBBYEFEj42q+MBZ7UjiNGCPO8TUxy5ox9MD8GA1Ud\nIwQ4MDaAFG52QCmlEWLXRZMQaBpGuecDHylCoROkETAPMQ0wCwYDVQQDEwR0ZXN0\nggkAxcaXrRjTYeAwEwYDVR0lBAwwCgYIKwYBBQUHAwIwCwYDVR0PBAQDAgeAMA0G\nCSqGSIb3DQEBCwUAA4IBgQDaUOQZ7AHMEpnJlz/vkcb6NSHzVVDOPbavWtwa0HkC\nHZTGIQVdMGLNHEPIhKjLZgdDsNLbNy/tVeaHTSp+OX+ZlfKLvrI7GrSyVF6THhUu\nQZ6iqK7lE7oWqEDPdZa+d8YaRm8ox3g7SsuqcWRS4cYFrxsGg++7tOhHx2kqWR4T\nB3Y1pN+TbrkmiofRuYrhgEYaGqZFbUyKwHWveEkuVq5uBL4ra61X4BT7XngIkAMe\nat/ugFEHl5FkFXF4444STn5lajisP1rG2y4ztiWE5Xgwl/Wkn2ozqIto8MeRK9WG\nkeEAnWdEGl2PDkzVnu5XbuPjhhk7d1Exjf6B4BN+AyAO6Pv0yhGTh4/hVkdhtMfQ\nz+Re5rff2p5zqI068UynVgXbZmyjFdzGF9x3731PADuODxq7rgfkp/fCpPUlMDjq\nlSqrt2kHQR8WrlpLZyYQgM4U9czHiYAOe0OZjXQvnttyKXVetCcFL5vEjnCUiBim\nzVbah7lMtSuacWEE5dAUDjQ=\n-----END CERTIFICATE-----\n", n.ClientCrt)
	th.AssertEquals(t, "-----BEGIN PRIVATE KEY-----\n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\njB8vl/PiudSrzCa2sKK6xZk=\n-----END PRIVATE KEY-----\n", n.ClientKey)
}