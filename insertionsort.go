package main

// InsertionSort performs an insertion sort in place on the specified array.
func InsertionSort(src []int) {
	// Start at 2nd value.
	for i := 1; i < len(src); i++ {
		// If this item is less than the item to our left, start shifting it left.
		if src[i] < src[i-1] {

			for j := i; j > 0; j-- {
				// Perform the swap.
				tmp := src[j-1]
				src[j-1] = src[j]
				src[j] = tmp

				// Finish swapping if our left value is <=
				// or we shifted all the way to the start.
				if j == 1 || src[j-2] <= src[j-1] {
					break
				}
			}

		}
	}
}
