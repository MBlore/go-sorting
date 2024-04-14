package main

import "fmt"

func main() {
	fmt.Println("Insertion Sort:")
	items := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(items)
	InsertionSort(items)
	fmt.Println(items)

	fmt.Println("Bubble Sort:")
	items = []int{90, 80, 70, 40, 50, 60, 30, 20, 10}
	fmt.Println(items)
	BubbleSort(items)
	fmt.Println(items)

	fmt.Println("Selection Sort:")
	items = []int{900, 800, 700, 400, 500, 600, 300, 200, 100}
	fmt.Println(items)
	SelectionSort(items)
	fmt.Println(items)

	fmt.Println("Merge Sort:")
	items = []int{9000, 8000, 7000, 4000, 5000, 6000, 3000, 2000, 1000}
	fmt.Println(items)
	MergeSort(items)
	fmt.Println(items)

	fmt.Println("Quick Sort:")
	items = []int{90000, 80000, 70000, 40000, 50000, 60000, 30000, 20000, 10000}
	fmt.Println(items)
	QuickSort(items)
	fmt.Println(items)
}
