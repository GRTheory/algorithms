package heapsort

type HeapStruct struct {
}

func DefaultHeapStruct() HeapStruct{
	return HeapStruct{}
}

func (h HeapStruct) Parent(i int) int {
	return i >> 1
}

func (h HeapStruct) Left(i int) int {
	return i << 1
}

func (h HeapStruct) Right(i int) int {
	return i<<1 + 1
}
