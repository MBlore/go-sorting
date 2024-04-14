package main

import (
	"fmt"
	"math/rand"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 10},
	{input: 100},
	{input: 1000},
	{input: 10000},
}

var tableLarge = []struct {
	input int
}{
	{input: 10},
	{input: 100},
	{input: 1000},
	{input: 10000},
	{input: 100000},
	{input: 1000000},
	{input: 10000000},
}

func makeArray(size int) []int {
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(size)
	}

	return numbers
}

func BenchmarkBubbleSort(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("Items %d", v.input), func(b *testing.B) {
			data := makeArray(v.input)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				BubbleSort(data)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("Items %d", v.input), func(b *testing.B) {
			data := makeArray(v.input)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SelectionSort(data)
			}
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("Items %d", v.input), func(b *testing.B) {
			data := makeArray(v.input)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				InsertionSort(data)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for _, v := range tableLarge {
		b.Run(fmt.Sprintf("Items %d", v.input), func(b *testing.B) {
			data := makeArray(v.input)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				MergeSort(data)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for _, v := range tableLarge {
		b.Run(fmt.Sprintf("Items %d", v.input), func(b *testing.B) {
			data := makeArray(v.input)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				QuickSort(data)
			}
		})
	}
}
