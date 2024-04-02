package main

import (
	"fmt"
)

func main() {
	nums := []int{7, 4, 5, 3, 1, 2, 6, 0}
	//counter := 0
	for i := len(nums); i > 0; i-- {
		for j :=0; j <  i; j++ {
			if nums[j] > nums[i-1] {
				nums[i-1], nums[j] = nums[j], nums[i-1]
				//counter++
			}
		}
	}
	fmt.Println(nums)
}
