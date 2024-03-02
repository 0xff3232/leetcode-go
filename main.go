// https://usaco.org/index.php?page=viewproblem2&cpid=966

package main

import (
	"fmt"
	"time"
)

func LinearMooBuzz(n int) int {

	r := []int{}
	i := 1
	for len(r) < n {
		if i%3 == 0 || i%5 == 0 {
			i++
			continue
		} else {
			r = append(r, i)
			i++
		}
	}
	fmt.Println(r)

	return r[len(r)-1]
}

func isValid(n int, x int) bool {
	return x-(x/3+x/5-x/15) >= n
}

func BMooBuzz(n int) int {
	left, right := 0, 10000000000

	for left < right {
		mid := (left + right) / 2
		if isValid(n, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	num := 100

	// LinearMooBuzz
	start := time.Now()
	fmt.Println("LinearMooBuzz:", LinearMooBuzz(num))
	fmt.Println("Time taken for LinearMooBuzz:", time.Since(start))

	// BMooBuzz
	start = time.Now()
	fmt.Println("BMooBuzz:", BMooBuzz(num))
	fmt.Println("Time taken for BMooBuzz:", time.Since(start))
}
