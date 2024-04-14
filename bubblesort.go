package main

// BubbleSort performs a bubble sort in place on the specified array.
func BubbleSort(src []int) {
	// First pass is to the far right of the array, decrementing per pass.
	for i := len(src) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if src[j] > src[j+1] {
				// Perform swap to shift the larger item over to the right.
				tmp := src[j+1]
				src[j+1] = src[j]
				src[j] = tmp
			}
		}
	}
}
