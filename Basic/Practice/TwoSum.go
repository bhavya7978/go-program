/*Given an array of integers nums and an integer target,
 return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution,
 and you may not use the same element twice.
You can return the answer in any order.*/
package main

import "fmt"

func main() {

	var target, l int
	fmt.Println("Enter the length of array:")
	fmt.Scan(&l)

	var arr = make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Enter the target sum:")
	fmt.Scan(&target)
	result := twoSum(arr, target)
	fmt.Println(result)
}
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[j]+nums[i] == target {

				return []int{i, j}
			}
		}
	}
	return nil
}
