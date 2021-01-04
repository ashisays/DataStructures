package dp

import (
    "fmt"
    "testing"
)

func TestLCS(t *testing.T) {
    fmt.Println("LCS is ",Lcs("Ashish","XXXXXXXXX",6,9))
    if Lcs("Ashish","XXXXXXXXX",6,9) == 0 {
      fmt.Println("The sequence is as expected.",)
    } else {
      t.Errorf("Failed ! got unexpected result !!")
    }
    fmt.Println("LCS is ",Lcs("Ashish","Ashish",6,6))
    if Lcs("Ashish","Ashish",6,6) == 6 {
      fmt.Println("The sequence is as expected.")
    } else {
      t.Errorf("Failed ! got unexpected result !!")
    }
}

func TestLCSDP(t *testing.T) {
    if LcsDP("Ashish","XXXXXXXXX",6,9) == 0 {
      fmt.Println("The sequence is as expected.")
    } else {
      t.Errorf("Failed ! got unexpected result !!")
    }
    fmt.Println("LCS is ",Lcs("Ashish","Ashish",6,6))
    if LcsDP("Ashish","Ashish",6,6) == 6 {
      fmt.Println("The sequence is as expected.")
    } else {
      t.Errorf("Failed ! got unexpected result !!")
    }
}