package main

import (
	"fmt"
	"lib"
	"os"
)

func main() {
	os.Exit(foo())
}

func foo() int {
	fmt.Printf("Hello World! 1 + 2 = %d\n", lib.SumInt(1, 2))
	return 0
}
