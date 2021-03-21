package heap

type Heap struct {
  HeapSize int
  Array []int
}

func Parent(i int) int {
  return i/2
}

func Left(i int) int {
  return 2*i+1
}

func Right(i int) int {
  return 2*i+2
}

func (heap *Heap)MaxHeapify(i int){
  l, r := 2*i+1, 2*i+2
	largest := i
  if l < heap.HeapSize && heap.Array[l]>heap.Array[i]{
    largest = l
  }
  if r < heap.HeapSize && heap.Array[r]>heap.Array[largest]{
    largest = r
  }
  if largest != i {
    heap.Array[i],heap.Array[largest] = heap.Array[largest],heap.Array[i]
    heap.MaxHeapify(largest)
  }
}

func MakeMaxHeap(array []int) Heap {
  heap := Heap{HeapSize : len(array), 
                Array: array}
  for i := len(array) / 2; i >= 0; i-- {
		heap.MaxHeapify(i)
	}
  return heap
}

func (heap *Heap) Show() []int {
  return heap.Array
}

func (heap *Heap) GetSize() int {
  return heap.HeapSize
}

func (heap *Heap) GetMaximum () int {
  return heap.Array[1]
}

func (heap *Heap) HeapExtractMax() (int,int) {
  if heap.HeapSize <1 {
    return 0,1
  }

  max := heap.Array[1]
  heap.Array[1] = heap.Array[heap.HeapSize-1] 
  heap.HeapSize--
  heap.MaxHeapify(1)
  return max,0
}

func (heap *Heap) IncreaseKey(i int,v int) int{
  if i< heap.HeapSize {
    heap.Array[i] =v
    for i>0 && heap.Array[Parent(i)] < heap.Array[i] {
      heap.Array[i],heap.Array[Parent(i)] = heap.Array[Parent(i)],heap.Array[i]
      i = Parent(i) 
    }   
  } else {
    return 1
  }
  return 0
}

func (heap *Heap) Insert(v int) {
  heap.HeapSize++
  heap.Array = append(heap.Array,v)
  heap.IncreaseKey(heap.HeapSize-1, v)
}