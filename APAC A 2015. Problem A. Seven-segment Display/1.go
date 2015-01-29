package main

import (
	"fmt"
	"os"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

var i2led []int = []int{0x7e, 0x30, 0x6d, 0x79, 0x33, 0x5b, 0x5f, 0x70, 0x7f, 0x7b}

func GetStep(x int, led int) (newgood int, newbad int) {
	newgood = x & led
	newbad = x ^ led
	return newgood, newbad
}

func GetResult(arr []int) string {
	res := -1
top:
	for x := 0; x < 10; x++ {
		good := 0
		bad := 0
		nextx := x
		for _, led := range arr {
			xled := i2led[nextx]
			newgood, newbad := GetStep(xled, led)
			good = good | newgood
			bad = bad | newbad

			// if x == 7 {
			// 	fmt.Printf("idx %v\n", idx)
			// 	fmt.Printf("xled %.7b\n", xled)
			// 	fmt.Printf("good %.7b\n", good)
			// 	fmt.Printf("bad %.7b\n", bad)
			// }

			if good&bad != 0 ||
				//				xled&newbad != 0 ||
				led&newbad != 0 {
				continue top
			}
			nextx = (10 + nextx - 1) % 10
		}

		xled := i2led[nextx] & (^bad)
		//fmt.Printf("going res %v\n", xled)
		if xled&good != xled ||
			(res >= 0 && res != xled) {
			return "ERROR!"
		}
		res = xled
	}
	if res >= 0 {
		return fmt.Sprintf("%.7b", res)
	} else {
		return "ERROR!"
	}

}

func main() {
	var cases int
	fmt.Scan(&cases)
	var arr []int
	for c := 0; c < cases; c++ {
		arr = []int{}
		var N int
		fmt.Scan(&N)
		for i := 0; i < N; i++ {
			var el int
			fmt.Scanf("%b", &el)
			arr = append(arr, el)
		}
		fmt.Printf("Case #%v: %v\n", c+1, GetResult(arr))

	}
	os.Exit(0)
}
