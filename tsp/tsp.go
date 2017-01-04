package main

import "fmt"

func main()  {
    var n int
    fmt.Println("Enter number of nodes:")
    fmt.Scanf("%d",&n)
    matrix := make([][]uint8, n) 
    for i,_ := range matrix {
        matrix[i] = make([]uint8, n+1)
        fmt.Println("Enter elements of row ",i)
        for k,_ := range matrix[i]{
            fmt.Scanf("%d",&matrix[i][k])
        }
    }
    fmt.Println(matrix)
}

func tsp(a []int) int  {
    var minimum = 999 
    var next int
    for i := 1; i < len(a); i++ {
        
    }
    return next
}

func next(a [][]uint8) int  {
    var next int
    var visited int
    for i := 1; i<= len(a)-1; i++{
        
    }
}
