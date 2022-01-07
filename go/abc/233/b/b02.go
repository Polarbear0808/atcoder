package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc, wr = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)

func scanString() string { sc.Scan(); return sc.Text() }
func scanRunes() []rune  { return []rune(scanString()) }
func scanInt() int       { a, _ := strconv.Atoi(scanString()); return a }
func scanInt64() int64   { a, _ := strconv.ParseInt(scanString(), 10, 64); return a }
func scanFloat() float64 { a, _ := strconv.ParseFloat(scanString(), 64); return a }
func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func print(S ...interface{}) {
	fmt.Fprint(wr, S...)
}
func println(S ...interface{}) {
	fmt.Fprintln(wr, S...)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1001), 1001001)
	L, R := scanInt(), scanInt()
	L--
	R--
	s := scanString()
	r := []rune(s[L : R-1])
	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}
	print(s[:L])
	print(string(r))
	println(s[R+1:])
}
