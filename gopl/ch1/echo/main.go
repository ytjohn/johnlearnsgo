// Echo1 prints its command-line arguments
package main

import (
	"os"
	"fmt"
)

func main() {
	var s, sep, cmd string
	cmd = os.Args[0]
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(cmd, s)
	fmt.Println(s)
}