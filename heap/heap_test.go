package heap;

import (
    "fmt"
    "testing"
)

func TestHeap(t *testing.T) {
  heap := MakeMaxHeap([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7})
  fmt.Println(heap.Show())
}


func BenchmarkHeap(b *testing.B) {
  for i:=0; i < b.N ; i++ {
    MakeMaxHeap([]int{27,17,3,16,13,10,1,5,7,12,4,8,9,0})
  }
}


func TestHeapSort(t *testing.T) {
  heap := HeapSort([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7})
  fmt.Println(heap)
}


func BenchmarkHeapSort(b *testing.B) {
  for i:=0; i < b.N ; i++ {
    HeapSort([]int{27,17,3,16,13,10,1,5,7,12,4,8,9,0})
  }
}
