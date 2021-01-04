package algo

import (
    "fmt"
    "testing"
)

func TestMaxPairwiseProduct(t *testing.T) {
    if MaxPairwiseProduct(1,2,3,4,5,6) == 30 {
      fmt.Println("The sequence is as expected.")
    } else {
      t.Errorf("Failed ! got unexpected result !!")
    }
    if MaxPairwiseProduct(1,5,3,4,2,6) == 30 {
      fmt.Println("The sequence is as expected.")
    } else {
      t.Errorf("Failed ! got unexpected result !!")
    }
    
    if MaxPairwiseProduct(1,6,3,4,2,
    ) == 36 {
      fmt.Println("The sequence is as expected.")
    } else {
      t.Errorf("Failed ! got unexpected result !!")
    }
}


