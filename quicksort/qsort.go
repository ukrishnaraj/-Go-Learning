package main

import (
    "fmt"
    "math/rand"
    )

/*func qsort(a []int)[]int  {
    if len(a)<2 {
        return a
    }
    left,right := 0,len(a)-1
    pivotIndex := rand.Int() % len(a)
    a[pivotIndex], a[right] = a[right], a[pivotIndex]
    for i,_ := range a {
        fmt.Println(i)
        if a[i] < a[right] {
        a[i], a[left] = a[left], a[i]
        left++
        }
    }
    a[left], a[right] = a[right], a[left]

    qsort(a[:left])
    qsort(a[left + 1:])
    return a
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

