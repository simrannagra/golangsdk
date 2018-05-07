/*
Package stacks enables management and retrieval of Stack Template
Stack service.


Example to Get Stack Template

result := stacktemplates.Get(client, "stackTrace", "e56cac00-463a-4e27-be14-abf414fc9816")
out, err := result.Extract()
fmt.Println(out)

*/

package stacktemplates
