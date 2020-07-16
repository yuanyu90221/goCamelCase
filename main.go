package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	fmt.Printf("%d\n", camelCaseSol(s))
}

func camelCaseSol(s string) int {
	cap := regexp.MustCompile(`[A-Z]`)
	ret := len(cap.Split(s, -1))
	return ret
}
