/*
Package Clusters enables management and retrieval of Clusters
Cluster service.

Example to List Clusters

	listOpts:=clusters.ListOpts{}
	allClusters,err1:=clusters.List(client,listOpts).Extract()
	if err != nil {
		panic(err)
	}

	for _, cluster := range allClusters {
		fmt.Printf("%+v\n", cluster)
	}

Example to Create a cluster
cluster:=clusters.CreateOpts{Kind:"Cluster",
								ApiVersion:"v3",
								Metadata:clusters.CreateMetaData{Name:"test-d12"},
								Spec:clusters.Spec{Type: "VirtualMachine",
													Flavor: "cce.s1.small",
													Version:"v1.7.3-r10",
													HostNetwork:clusters.HostNetwokSpec{
																					VpcId:{vpc_id}
																					SubnetId:{subnet_id},},
													ContainerNetwork:clusters.ContainerNetworkSpec{Mode:"overlay_l2"},
													},

	}
 	status,err1:=clusters.Create(client,cluster).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a cluster

	cluster:=clusters.UpdateOpts{clusters.UpdateSpec{"test"}}

	out,err:=clusters.Update(client,{cluster_id},cluster).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a cluster

	del,err:=clusters.Delete(client,{cluster_id}).Extract()
	if err != nil {
		panic(err)
	}
*/
package clusters
