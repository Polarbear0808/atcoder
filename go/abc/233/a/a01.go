package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	ans := (y - x) / 10

	if y-x > 0 {
		if (y-x)%10 != 0 {
			ans += 1
		}
	} else {
		ans = 0
	}

	fmt.Println(ans)
}
