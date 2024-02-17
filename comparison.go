package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "123456789"
	s2 := "123456789"

	v := strings.EqualFold(s1, s2)

	fmt.Println(v)
}
