// Exercise 1.2: Modify the echo program to print the index and value of each
// of its arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("os.Args[%v]: %v\n", i, os.Args[i])
	}
}
