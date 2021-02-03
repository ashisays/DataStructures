package list

import (
    "fmt"
    "testing"
)


func TestShow(t *testing.T) {
    head := Node{val:1}
    for i:=1; i < 5; i++{
      head.insert(i+1)
    }
    head.show()
}

func TestKthElement(t *testing.T) {
    head := Node{val:1}
    for i:=1; i < 5; i++{
      head.insert(i+1)
    }
    find_kth_elem(&head,2)
}

func TestIntersection(t *testing.T) {
    head := Node{val:1}
    head1 := Node{val:10}
    for i:=1; i < 5; i++{
      head.insert(i+1)
      head1.insert(10-i)
    }
    head1.append(&head)
    head1.show()
    fmt.Printf("The interscection node is  :%v",find_intersection(&head,&head1))
}


func TestLoop(t *testing.T) {
    head := Node{val:1}
    for i:=1; i < 5; i++{
      head.insert(i+1)
    }
    fmt.Printf("The loop in  node is  :%v\n",head.find_cycle())
    head.append(&head)
    fmt.Printf("The loop in  node is  :%v",head.find_cycle())
}


