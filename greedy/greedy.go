package greedy

import (
  "sort"
  "fmt"
)

/*
Money Change Problem
Compute the minimum number of coins needed
to change the given value into coins with denominations
1, 5, and 10.
Input: Integer money.
Output: The minimum number
of coins with denominations 1, 5,
and 10 that changes money
*/
func MoneyChange(m int) int {
  numCoins := 0
  for m > 0{
    if m >= 10{
      m = m -10
      numCoins++
    } else if m >= 5 {
      m = m - 5
      numCoins++
    } else {
      m--
      numCoins++
    }
  }
  return numCoins
}

/*
Maximizing the Value of the Loot Problem
Find the maximal value of items that fit into the
backpack.
Input: The capacity of a backpack
W as well as the weights
(w1; : : : ;wn) and per pound prices
(p1; : : : ;pn) of n different compounds.
Output: The maximum total price
of items that fit into the backpack
of the given capacity: i.e., the maximum
value of p1  u1 +    + pn  un
such that u1 +    + un  W and 0 
ui  wi for all i.
*/
type iVal struct {
  w float64
  p float64
  r float64
}
func MaxLoot(w float64,items []float64,prices []float64)  string {
  p := 0.0
  ratio := make([] iVal,0) 
  for i,val :=range items {
      ratio = append(ratio,iVal{val,prices[i],prices[i]/val})
  }
  
  sort.Slice(ratio, func(i, j int) bool {
  return ratio[i].r > ratio[j].r
  })
  for _,ival := range ratio{
    if w == 0 {
      break
    } else if ival.w <= w{
      w = w - ival.w
      p = p + ival.p
    } else {
      p = p+(ival.w-w)*ival.r
      w = 0
    }
  }
  s := fmt.Sprintf("%.2f", p)
  return s
}