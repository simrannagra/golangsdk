// Package stacks provides operation for working with Heat stacks. A stack is a
// group of resources (servers, load balancers, databases, and so forth)
// combined to fulfill a useful purpose. Based on a template, Heat orchestration
// engine creates an instantiated set of resources (a stack) to run the
// application framework or component specified (in the template). A stack is a
// running instance of a template. The result of creating a stack is a deployment
// of the application framework or component.

/*
Package resources enables management and retrieval of
RTS service.

Example to List Resources

lis:=stacks.ListOpts{SortDir:stacks.SortAsc,SortKey:stacks.SortStatus}
	getstack,err:=stacks.List(client,lis).AllPages()
	stacks,err:=stacks.ExtractStacks(getstack)
	fmt.Println(stacks)
*/
package stacks
