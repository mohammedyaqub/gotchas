//Accessing Non-Existing Map Keys
package main
import "fmt"
/*Checking for the appropriate "zero value" can be used to determine 
if the map record exists, but it's not always reliable
 (e.g., what do you do if you have a map of booleans where the "zero value" is false).
  The most reliable way to know if a given map record exists is to check the second 
  value returned by the map access operation.
  */
func main() {
	//wrong way doing 
	// 	m:=map[string]string{"one":"1","two":"2","three":""}
	// if v:=m["three"]; v=="" {
	// 	fmt.Println("no entry")
	// 	}

	m:=map[string]string{"one":"1","two":"2","three":""}

	if _,ok:=m["three"];ok{
		fmt.Println("no entry with value")
	}

	if v,ok:=m["two"];ok{
		fmt.Println("entry with value is ",v)
	}
}