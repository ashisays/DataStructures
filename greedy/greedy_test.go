package greedy

import (
    "fmt"
    "testing"
)


func TestMoneyChange(t *testing.T) {
    fmt.Println("Money CHange is ",MoneyChange(2))
    if MoneyChange(0) == 0 {
       fmt.Println("The Coin used is as expected.",)
    } else {
       t.Errorf("Failed ! got unexpected result !!")
    }

    fmt.Println("LCS is ",MoneyChange(4))
    if MoneyChange(4) == 4 {
      fmt.Println("The sequence is as expected.")
    } else {
      t.Errorf("Failed ! got unexpected result !!")
    }
}


func TestMaxLoot(t *testing.T) {
    w := []float64{1,3,4,5,6,6}
    p := []float64{10,30,50,60,70,90}
    fmt.Println("Max loot  is ",MaxLoot(20,w,p))
}