package main

// Global temp arrays - not thread safe, but fine for testing purposes.
// Shows that a merge sort has a cost of a small amount of memory allocation to perform where simpler
// sorts are in-place and don't need any extra memory.
var la []int
var ra []int

func msmerge(src []int, l int, m int, r int) {
	// Get sizes of the two arrays to be merged.
	n1 := m - l + 1
	n2 := r - m

	// Alloc storage arrays.
	if len(la) < n1 {
		la = make([]int, n1)
	}
	if len(ra) < n2 {
		ra = make([]int, n2)
	}

	i, j := 0, 0

	// Copy data to temp arrays.
	for i = 0; i < n1; i++ {
		la[i] = src[l+i]
	}
	for j = 0; j < n2; j++ {
		ra[j] = src[m+1+j]
	}

	// Merge the temp arrays.
	i, j = 0, 0
	k := l

	// Initial index array.
	for i < n1 && j < n2 {
		if la[i] <= ra[j] {
			src[k] = la[i]
			i++
		} else {
			src[k] = ra[j]
			j++
		}

		k++
	}

	// Copy remaining elements from la.
	for i < n1 {
		src[k] = la[i]
		i++
		k++
	}

	// Copy remaining elements from ra.
	for j < n2 {
		src[k] = ra[j]
		j++
		k++
	}
}

func mssort(src []int, l int, r int) {
	if l < r {
		// Find mid point.
		m := l + (r-l)/2

		// Sort the halves.
		mssort(src, l, m)
		mssort(src, m+1, r)

		// Merge the sorted halves.
		msmerge(src, l, m, r)
	}
}

// MergeSort performs a merge sort in place on the specified array.
func MergeSort(src []int) {
	mssort(src, 0, len(src)-1)
}
