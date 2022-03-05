package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	a := append(nums1, nums2...)
	fmt.Println(a)

	for i := 0; i <= len(a)-1; i++ {
		for j := i + 1; j <= len(a)-1; j++ {

			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}

		}
	}
	fmt.Println(a)
	if len(a)%2 == 0 {
		median := (a[(len(a)/2)-1] + a[len(a)/2])
		var final float64
		final = float64(median) / 2.0
		return final
	} else {
		result := a[(len(a)-1)/2]
		return float64(result)

	}

}
