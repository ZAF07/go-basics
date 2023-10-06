package main

import (
	"fmt"
	"math"
)

/* Math Abs
The Abs function returns a single value of type float64. This value represents the absolute value of the given argument. So, all returned values are greater or equal to 0.
*/

/* Task
Find the lowest possible number in this array. If the lowest number is a negative number and there is a positive equivalent to it, return the positive number

Example: [-1, 4, 2, 1]
Result: 1 (-1 is technically the lowest possible number but we want a positive number if possible)
*/

func main() {
	nums := []int{1, 3, -1, 2, 1, 5}

	fmt.Println(closestToZero(nums))
}

func closestToZero(nums []int) int {

	result := nums[0]
	if len(nums) < 1 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		val := nums[i]

		// use the absolute value of the integers for value checks here (math.Abs returns the positive value of the given input. so given -1, 1 will be returned)
		if int(math.Abs(float64(val))) <= int(math.Abs(float64(result))) {
			// check if the actual value of the curret value is a negative, if so, continue
			if val < 0 {
				continue
			}
			result = val
		}
	}
	return result
}
