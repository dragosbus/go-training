package main

import "fmt"

//last digit
func strings(x int) string {
	//coerce int to a string
	xStr := fmt.Sprintf("%d", x)
	length := len(xStr)
	return xStr[length-1:]
}
