package stringds

func UniqueChar(str string) bool {
  var checker int64
  for _, c := range str {
    bitIndex := c - 'a'
    if (checker & (1 << bitIndex)) > 0 {
      return false
    }
    checker = checker | (1 << bitIndex)
  }
  return true
}