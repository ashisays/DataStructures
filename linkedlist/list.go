package list

import "fmt"

type Node struct{
  val int
  next *Node
}

func (head *Node) show(){
  temp := head
  fmt.Printf("%v",temp.val)
  for temp.next != nil{
    temp = temp.next
    fmt.Printf("-->%v",temp.val)
  }
  fmt.Println("")
}

func (head *Node) insert(val int){
  temp := head
  for temp.next != nil{
    temp = temp.next
  }
  temp.next = &Node{val:val}
}

func (head *Node) len() int {
  temp := head
  i :=0
  for temp.next != nil{
    temp = temp.next
    i++
  }
  return i
}

func (head *Node) append(head1 *Node) {
  temp := head
  for temp.next != nil{
    temp = temp.next
  }
  temp.next = head1
}

func find_kth_elem(head *Node,k int) int {
  temp := head
  if temp.next == nil{
    return 1
  } else {
    i := find_kth_elem(temp.next,k)+1
    if i == k {
      fmt.Printf("%v element is : %v\n",k,temp.val)
    }
    return i
  }    
}

func find_intersection(head1 *Node,head2 *Node) int {
  len1 := head1.len()
  len2 := head2.len()
  fmt.Printf("Len1 : %v, Len2: %v",len1,len2)
  if len1 > len2{
    for i:=0;i<(len1-len2);i++{
      fmt.Printf("%v-->",head1.val)
      head1 = head1.next
    }
  } else if len2 > len1{
    for i:=0;i<(len2-len1);i++{
      head2 = head2.next
    }
  }
  fmt.Printf("\nll1 %v-->",head1.val)
  fmt.Printf("\nll2 %v-->",head2.val)
  for head1.next != nil && head2.next != nil{
    if head1 == head2{
      return head1.val
    } else{
      head1 = head1.next
      head2 = head2.next
    }
  }
  return -1
}
