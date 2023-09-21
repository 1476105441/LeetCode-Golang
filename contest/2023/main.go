package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 35, 532, 53, 54, 12, 3}

	b := make([]int, len(a))

	copy(b, a)

	sort.Ints(b)

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(1, 2)
}
