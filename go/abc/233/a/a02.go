package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var ans int
	if x >= y {
		ans = 0
	} else {
		ans = (y - x) / 10

		if (y-x)%10 > 0 {
			ans += 1
		}
	}

	fmt.Println(ans)
}
