// This is a package
package main

import (
	"fmt"
	"runtime"
)

/*
Main package
Multiline comment
*/
func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			hey()
		} else {
			bye()
		}
	}
	fmt.Println(runtime.NumCPU())
}
