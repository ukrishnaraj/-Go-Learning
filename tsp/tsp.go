package main

import "fmt"

var cost int
func main()  {
    var n,start int
    fmt.Println("Enter number of nodes:")
    fmt.Scanf("%d\n",&n)
    matrix := make([][]int, n) 
    for i,_ := range matrix {
        matrix[i] = make([]int, n)
        fmt.Println("Enter elements of row ",i)
        for k,_ := range matrix[i]{
            fmt.Scanf("%d\n",&matrix[i][k])
        }
    }
    fmt.Println("Enter the starting node :")
    fmt.Scanf("%d\n",&start)
    fmt.Println(tsp(matrix,start))
}

func tsp(a [][]int,start int) []int  {
    travel := make([]int, len(a)+1)
    visited := make([]int, len(a))
    for i := range visited{
        visited[i]=0
    }
    for i:=0; i< len(a)-1; i++{
        travel[i] = start
        if i==len(a)-1 {
            travel[i+1]= travel[0]
        }
        visited[start] = 1 
        new := nextNode(a[start],visited,start)
        if new !=0{
            start = new
        } else {
            start = travel[0]
        }
    }
    return travel
}

func nextNode(a,visited []int,x int) int  {
    var next int
    min := 999
    for i := 0; i< len(a); i++{ 
        
        if a[i] != 0 && visited[i]==0 && a[i]<min{
                min = a[i]
                next = i
        }
    }
    fmt.Println(".$$52##",min)
    return next
}
