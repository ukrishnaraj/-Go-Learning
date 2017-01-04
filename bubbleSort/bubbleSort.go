package main

import (
    "fmt"
    )

func bubbleSort(unsorted []int)  {
    if len(unsorted) < 2 {
        return 
    }
    var n int
    n = len(unsorted)
    for i := 0; i < n; i++ {
        for j := n-1; j >= i+1; j--{
            if unsorted[j]<unsorted[j-1] {
                unsorted[j],unsorted[j-1] = unsorted[j-1],unsorted[j]
            }
        }
    }
}

func main()  {
    var count int
    fmt.Println("Enter count:")
    fmt.Scanf("%d",&count)
    toSort := make([]int,count+1)
    fmt.Println("Enter the numbers:")
    for k,_ := range toSort{
        fmt.Scanf("%d",&toSort[k])
    }
    fmt.Println("Unsorted:",toSort[1:count+1])
    bubbleSort(toSort[1:count+1])
    fmt.Println("After sorting:",toSort[1:count+1])
}