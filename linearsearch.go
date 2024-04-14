package main

// BinarySearch used a binary search algorithm to find 'val' in the array 'src'.
func LinearSearch(src []int, val int) int {
	for i := 0; i < len(src); i++ {
		if src[i] == val {
			return i
		}
	}

	return -1
}
