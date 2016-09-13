package main

import "fmt"

func writeMsg(a, b int) {
	fmt.Print(a + b)
}

func main() {
	// call the function
	writeMsg(1, 2)
}
