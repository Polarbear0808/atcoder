package main

import "fmt"

func main() {
	s := "あいうえお😂"
	for i := 0; i < len(s); i++ {
		b := s[i]
		fmt.Printf("%c ", b)
	}
	fmt.Println("")

	for _, r := range s {
		fmt.Printf("%c", r)
	}
	fmt.Println("")

	for _, r := range []rune(s) {
		fmt.Printf("%c", r)
	}
	fmt.Println("")

	for i := 0; i < len(s); i++ {
		b := s[i]
		fmt.Println(b)
	}
	for _, r := range s {
		fmt.Println(r)
	}

	for _, r := range []rune(s) {
		fmt.Println(r)
	}
}
