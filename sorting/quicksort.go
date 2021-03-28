package sorting

func QuickSort(a []int) []int {
    if len(a) < 2 {
        return a
    }
     
    start, end := 0, len(a)-1     
    for i, _ := range a {
        if a[i] < a[end] {
            a[start], a[i] = a[i], a[start]
            start++
        }
    }
     
    a[start], a[end] = a[end], a[start]
     
    QuickSort(a[:start])
    QuickSort(a[start+1:])
     
    return a
}