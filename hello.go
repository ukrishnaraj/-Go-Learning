package main

import "fmt"

func main() {
    var i,j,k int
    fmt.Printf("enter numbers:")
    fmt.Scanf("%d ",&i)
    fmt.Scanf("%d ",&j)
    k = i + j
    fmt.Printf("sum is %d",k)
}