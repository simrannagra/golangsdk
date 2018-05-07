/*
Package stacks enables management and retrieval of Stacks
Stack service.


Example to List Stacks

lis := stacks.ListOpts{SortDir:stacks.SortAsc,SortKey:stacks.SortStatus}
allStacks, err := stacks.List(client,lis).AllPages()
stacks,err:=stacks.ExtractStacks(allStacks)
fmt.Println(stacks)
*/
package stacks

