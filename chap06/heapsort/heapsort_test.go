package heapsort_test

import (
	"testing"

	"github.com/GRTheory/algorithms/chap06/heapsort"
)

type slice []int

// Len is the number of elements in the collection.
func (s slice) Len() int {
	return len(s)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (s slice) Less(i int, j int) bool {
	return s[i] < s[j]
}

// Swap swaps the elements with indexes i and j.
func (s slice) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s slice) Subset(n int) any {
	return s[:n]
}

func TestGenerateMaxHeap(t *testing.T) {
	sort := heapsort.HeapSort{heapsort.HeapStruct{}}
	slices := slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sort.BuildMaxHeap(slices)
	t.Log(slices)
}

func TestHeapsort(t *testing.T) {
	sort := heapsort.HeapSort{heapsort.HeapStruct{}}
	slices := slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sort.HeapSort(slices)
	t.Log(slices)
}
