package main

import "fmt"

func qsort(arr []int)[]int  {
    size := len(arr)

    if size<2 {
        return arr
    }
    p := size % 2
    left := make([]int, 0, size)
    middle := make([]int, 0, size)
    right := make([]int, 0,size)
    for i := 0; i < size; i++ {
        if arr[i]<arr[p] {
            left = append(left,arr[i])
        }else if arr[i]==arr[p] {
            middle = append(middle,arr[i])
        }else {
            right = append(right,arr[i])
        }
    }
    left = qsort(left)
    right = qsort(right)

    left = append(left,middle...)
    left = append(left,right...)

    return left
}

func main()  {
    var n int
    fmt.Println("Enter Count:")
    fmt.Scanf("%d",&n)
    toSort := make([]int, n+1)
    fmt.Println("Enter the numbers:")
    for k,_ := range toSort{
        fmt.Scanf("%d",&toSort[k])
    }
    fmt.Println("Unsorted Array:",toSort[1:n+1])
    qsort(toSort[1:n+1])
    fmt.Println("Sorted Array:",toSort[1:n+1])
}