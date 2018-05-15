package testing

import (
	"github.com/huaweicloud/golangsdk/openstack/cce/v1/clusters"
)

const GetOutput = `
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
}`

var GetExpected = &clusters.RetrievedCluster{
	Kind:       "cluster",
	ApiVersion: "v1",
	Metadata: clusters.MetaData{Name: "terraform-test",
		ID:      "d83af16b-febd-4e52-bfb0-20850072e2cd",
		SpaceID: "80f4acb2-6403-0c6a-1cf2-07d02ea708cc"},
	Clusterspec: clusters.ClusterlistSpec{AZ: "eu-de-01",
		CPU:              1,
		Memory:           4096,
		VPC:              "vpc-terraform",
		VpcId:            "5232f396-d6cc-4a81-8de3-afd7a7ecdfd8",
		Subnet:           "subnet-28a5",
		Endpoint:         "https://192.168.0.45:5443",
		ExternalEndpoint: "https://160.44.197.88:5443",
		ClusterType:      "Single",
		SecurityGroupId:  "01042a5b-eef7-46a4-9076-4d35de563f3a",
		ClusterHostList: clusters.ClusterHostList{Kind: "list",
			ApiVersion: "v1",
			Metadata:   clusters.MetaData{},
			HostListSpec: clusters.HostListSpec{HostList: []clusters.Host{{Kind: "host",
				ApiVersion: "v1",
				Metadata: clusters.MetaData{Name: "terraform-test-node-1",
					ID:      "a535f227-e0fb-402f-94f8-d97354a8254d",
					SpaceID: "80f4acb2-6403-0c6a-1cf2-07d02ea708cc",
				},
				Hostspec: clusters.HostSpec{ClusterUuid: "d83af16b-febd-4e52-bfb0-20850072e2cd",
					ClusterName: "terraform-test",
					PrivateIp:   "192.168.0.53",
					PublicIp:    "160.44.194.122",
					Flavor:      "s1.medium",
					CPU:         1,
					Memory:      4096,
					HostId:      "a535f227-e0fb-402f-94f8-d97354a8254d",
					AZ:          "eu-de-02",
					SshKey:      "KeyPair-terraform",
					Volume:      []clusters.Volume{{DiskType: "root", DiskSize: 40, VolumeType: "SATA"}},
					Status: clusters.HostStatus{Capacity: clusters.Capacity{CPU: "1", Memory: "3844320Ki", Pods: "110"},
						Allocatable:     clusters.Capacity{CPU: "0.452", Memory: "3317984Ki", Pods: "110"},
						Conditions:      []clusters.Conditions{{Type: "OutOfDisk", Status: "False", Reason: "KubeletHasSufficientDisk"}},
						Addresses:       []clusters.Addresses{{Type: "InternalIP", Address: "192.168.0.53"}},
						DaemonEndpoints: clusters.DaemonEndpoints{KubeletEndpoint: clusters.KubeletEndpoint{Port: 10250}},
						NodeInfo: clusters.NodeInfo{MachineID: "f21bcad2f362498d924ac24efcdd3e2a",
							SystemUUID:              "69646393-240E-4B63-86DE-1DE5E341F732",
							BootID:                  "cae34bd3-277f-4a08-b429-6337b1bca8a0",
							KernelVersion:           "3.10.0-229.42.1.97.x86_64",
							OsImage:                 "EulerOS V2.0SP1",
							ContainerRuntimeVersion: "docker://1.11.2",
							KubeletVersion:          "cce-1.7.3.17+8857de1a267f8f-dirty",
							KubeProxyVersion:        "cce-1.7.3.17+8857de1a267f8f-dirty"},
						Images: []clusters.Images{{Names: []string{"fluentd:1.11"}, SizeBytes: 617825114}}}},
				Replicas: 1,
				Status:   "ACTIVE",
				Message:  "terraform-test-node-1",
			}},
			},
		},
	},

	ClusterStatus: clusters.ClusterStatus{Status: "AVAILABLE"},
	K8sVersion:    "1.7.3",
}
