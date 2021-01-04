/*Maximum Pairwise Product Problem
Find the maximum product of two distinct numbers
in a sequence of non-negative integers.
Input: A sequence of non-negative
integers.
Output: The maximum value that
can be obtained by multiplying
two different elements from the sequence.
Input format. The first line contains an integer n. The next line contains
n non-negative integers a1; : : : ;an (separated by spaces).
Output format. The maximum pairwise product.
Constraints. 2  n  2  105; 0  a1; : : : ;an  2  105.
*/
package algo

func MaxPairwiseProduct(a ...int) int {
  la := len(a)
  max := make([]int,2,2)
  for i:=0; i < la; i ++{
    if a[i] > max[0]{
      max[0] = i
    }
  }
  for i:=0; i < la; i ++{
    if a[i] > max[1] && i!= max[0]{
      max[1] = a[i]
    }
  }
  return max[0]*max[1]
}