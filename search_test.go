package main

import (
	"fmt"
	"testing"
)

var tableSearch = []struct {
	input int
}{
	{input: 10000000},
	{input: 100000000},
	{input: 1000000000},
}

func makeArrayOrdered(size int) []int {
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = i
	}

	return numbers
}

func BenchmarkBinarySearch(b *testing.B) {
	for _, v := range tableSearch {
		b.Run(fmt.Sprintf("Items %d", v.input), func(b *testing.B) {
			data := makeArrayOrdered(v.input)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				BinarySearch(data, data[len(data)/4*3])
			}
		})
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	for _, v := range tableSearch {
		b.Run(fmt.Sprintf("Items %d", v.input), func(b *testing.B) {
			data := makeArrayOrdered(v.input)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				LinearSearch(data, data[len(data)/4*3])
			}
		})
	}
}
