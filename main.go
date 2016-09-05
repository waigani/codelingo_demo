package main

import "fmt"

func writeMsg(a, b, c, d int) {
	fmt.Print(a+b+c+d)
}

func main() {
	// call the function
	writeMsg(1,2,3,4)
}