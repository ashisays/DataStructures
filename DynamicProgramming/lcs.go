package dp

func max(a,b int) int {
  if a > b {
    return a
  }
  return b
}

func Lcs(str1, str2 string, m , n int) int {
  if m == 0 || n == 0{
    return 0
  }
  if str1[m-1] == str2[n-1]{
    return 1 + Lcs(str1,str2,m-1,n-1)
  } else {
    return max(Lcs(str1,str2,m-1,n),Lcs(str1,str2,m,n-1))
  }
}

func LcsDP(str1, str2 string, m , n int) int {
  tab := make([][]int,m+1)
  for i := range tab {
        tab[i] = make([]int, n+1)
  }
  i, j := 0, 0
  for i = 0; i <= m; i++ {
      for j = 0; j <= n; j++ {
        if i == 0 || j == 0{
            tab[i][j] = 0
        }else if str1[i-1] == str2[j-1] {
                tab[i][j] = tab[i-1][j-1] + 1
        } else {
                tab[i][j] = max(tab[i-1][j], tab[i][j-1])
        }
      }
  }
  return tab[m][n]
}