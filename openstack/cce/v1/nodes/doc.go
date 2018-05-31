package nodes
/*

Package nodes enables management and retrieval of nodes
CCE service.
Example to List nodes
	listNodes := nodes.ListOpts{}
	clusterID := "d83af16b-febd-4e52-bfb0-20850072e2cd"
	allNodes, err := nodes.List(client,clusterID ).ExtractNode(listNodes)
	if err != nil {
		panic(err)
	}
	for _, node := range allNodes {
		fmt.Printf("%+v\n", node)
	}

 */