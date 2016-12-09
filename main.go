package main

import "fmt"

func writeAMsg(a, b int) {
	fmt.Print(a + b)
}

func main() {
	// call the function
	writeAMsg(1, 2)
}
