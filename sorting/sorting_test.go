package sorting

import (
    "fmt"
    "testing"
    "math/rand"
    "time"
)

func generateRandomSlice(size int) []int {
 
    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}

func TestUnique(t *testing.T) {
  a := [] int {2,4,5,1,3,7,8,9,6}
  b := QuickSort(a)
  fmt.Println("Sort Successful !!",b)
}

func BenchmarkQuickSort(b *testing.B) {
  for i:=0; i < b.N ; i++ {
    QuickSort(generateRandomSlice(20))    
  }
}