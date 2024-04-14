package main

func bsearch(src []int, val int, low int, high int) int {
	for low <= high {
		mid := low + (high-low)/2
		if src[mid] == val {
			// Found.
			return mid
		} else if src[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// Not found.
	return -1
}

// BinarySearch uses a binary search algorithm to find 'val' in the array 'src'.
func BinarySearch(src []int, val int) int {
	return bsearch(src, val, 0, len(src)-1)
}
