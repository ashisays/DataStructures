package heap

func HeapSort(array []int) []int {
  heap := MakeMaxHeap(array)
  for i := heap.HeapSize-1 ; i > 0; i -- {
    heap.Array[0],heap.Array[i] = heap.Array[i],heap.Array[0]
    heap.HeapSize--
    heap.MaxHeapify(0)
  } 
  return heap.Show()
}