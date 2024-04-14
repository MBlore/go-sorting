package main

// SelectionSort performs a selection sort in place on the specified array.
func SelectionSort(src []int) {
	for i := 0; i < len(src)-1; i++ {
		min := src[i]
		swap := -1

		// Find the smallest value for position i.
		for j := i + 1; j < len(src); j++ {
			if src[j] < min {
				min = src[j]
				swap = j
			}
		}

		// If we found a smaller value, swap them.
		if swap != -1 {
			tmp := src[i]
			src[i] = src[swap]
			src[swap] = tmp
		}
	}
}
