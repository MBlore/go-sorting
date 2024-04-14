package main

func qspartition(src []int, low int, high int) int {
	pivot := src[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if src[j] < pivot {
			i++

			// Swap.
			temp := src[i]
			src[i] = src[j]
			src[j] = temp
		}
	}

	// Swap the pivot.
	temp := src[i+1]
	src[i+1] = src[high]
	src[high] = temp

	return i + 1
}

func qssort(src []int, low int, high int) {
	if low < high {
		// Get a partitoning index.
		pi := qspartition(src, low, high)

		// Sort items before and after partition.
		qssort(src, low, pi-1)
		qssort(src, pi+1, high)
	}
}

// SelectionSort performs a selection sort in place on the specified array.
func QuickSort(src []int) {
	qssort(src, 0, len(src)-1)
}
