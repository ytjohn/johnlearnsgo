// Echo1 prints its command-line arguments
package main

import (
	"os"
	"fmt"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i)
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}