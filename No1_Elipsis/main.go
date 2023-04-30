package main

import (
	"fmt"
)

func main() {

	s1 := "Hello "
	s2 := "world!"
	s3 := " How are you?"
	result := concatStrings(s1, s2, s3)
	fmt.Printf(result)
}

func concatStrings(strs ...string) string {
	var result string

	for _, str := range strs {
		result += str
	}

	return result
}
