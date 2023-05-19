// the example for how to use PriorityDeque
package main

import (
	"cncsmonster/prioritydeque"
	"fmt"
)

func main() {
	pdq := prioritydeque.New(func(a, b any) bool {
		return a.(int) < b.(int)
	})
	arr := []int{8, 3, 1, 5, 2}

	for _, v := range arr {
		pdq.Push(v)
		fmt.Println("max:", pdq.Max(), "min:", pdq.Min())
	}
	fmt.Println("pop:", pdq.PopMax())
	fmt.Println("max:", pdq.Max(), "min:", pdq.Min())
	fmt.Println("pop:", pdq.PopMin())
	fmt.Println("max:", pdq.Max(), "min:", pdq.Min())
}
