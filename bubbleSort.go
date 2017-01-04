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
    //unsort := []int{33,1,5,92,63,12}
    var count int
    fmt.Println("Enter count:")
    fmt.Scanf("%d",&count)
    toSort := make([]int,count)
    fmt.Println("Enter the numbers:")
    for k,_ := range toSort{
        fmt.Scanf("%d",&toSort[k])
    }
    fmt.Println("Unsorted:",toSort)
    bubbleSort(toSort)
    fmt.Println("After sorting:",toSort)
}