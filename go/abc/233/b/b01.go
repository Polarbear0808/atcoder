package main

import (
	"fmt"
	"strings"
)

func main() {
	// get input (string array)
	var l, r int
	var s string
	var b strings.Builder
	fmt.Scanln(&l, &r)
	fmt.Scanln(&s)

	// build string array
	// left
	if l > 1 {
		for i := 0; i < l-1; i++ {
			fmt.Fprintf(&b, "%c", s[i])
		}
	}

	// center
	for i := r - 1; i > l-2; i-- {
		fmt.Fprintf(&b, "%c", s[i])
	}

	// right
	if r < len(s) {
		for i := r; i < len(s); i++ {
			fmt.Fprintf(&b, "%c", s[i])
		}
	}

	// output
	fmt.Println(b.String())
}
