package heapsort

import (
	"sort"
)

type Interface interface {
	Parent(i int) int
	Left(i int) int
	Right(i int) int
}

type MySortInterface interface {
	sort.Interface
	Subset(n int) any
}

type HeapSort struct {
	Interface
}

func (h HeapSort) HeapSort(slice MySortInterface) {
	h.BuildMaxHeap(slice)
	for i := slice.Len(); i >= 1; i-- {
		slice.Swap(0, i-1)
		temp := slice.Subset(i-1)
		h.MaxHeap(temp.(MySortInterface), 1)
	}
}

func (h HeapSort) BuildMaxHeap(slice MySortInterface) {
	for i := slice.Len() / 2; i > 0; i-- {
		h.MaxHeap(slice, i)
	}
}

func (h HeapSort) MaxHeap(slice MySortInterface, i int) {
	l := h.Left(i)
	r := h.Right(i)
	largest := i

	if l <= slice.Len() && slice.Less(largest-1, l-1) {
		largest = l
	}

	if r <= slice.Len() && slice.Less(largest-1, r-1) {
		largest = r
	}

	if largest != i {
		slice.Swap(i-1, largest-1)
		h.MaxHeap(slice, largest)
	}
}
