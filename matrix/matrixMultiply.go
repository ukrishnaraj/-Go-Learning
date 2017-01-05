package main

import "fmt"

func main()  {
    var ar,ac,br,bc int
    fmt.Println("(Matrix#1)Enter number of rows :")
    fmt.Scanf("%d\n",&ar)
    fmt.Printf("(Matrix#1)Enter number of columns:")
    fmt.Scanf("%d\n",&ac)
    a:= make([][]int, ar)
    for i,_:= range a{
        a[i] = make([]int, ac)
        fmt.Println("(Matrix#1)Row ",i,":")
        for j:= range a[i]{
            fmt.Scanf("%d\n",&a[i][j])
        }
    }
    fmt.Printf("(Matrix#2)Enter number of rows: ")
    fmt.Scanf("%d\n",&br)
    fmt.Printf("(Matrix#2)Enter number of columns: ")
    fmt.Scanf("%d\n",&bc)
    b:= make([][]int, br)
    for k:= range b{
        b[k] = make([]int, bc)
        fmt.Println("(Matrix#1)Row ",k,":")
        for m:= range b[k]{
            fmt.Scanf("%d\n",&b[k][m])
        }
    }
    product := multiply(a,b)
    fmt.Println(product)
}
func getColumn(array [][]int,x int) []int  {
    column := make([]int, 0)
    for t,_ := range array{
        column = append(column,array[t][x])
    }
    return column
}
func result(x,y [][]int) [][]int {
    z := make([][]int, len(x))
    for i := range z{
        z[i] = make([]int, len(getColumn(y,0)))
        for k := range z[i]{
            z[i][k] = 0
        }
    }
    return z
}

func multiply(x,y [][]int)[][]int  {
    p := result(x,y)
    fmt.Println(p)
    if len(x)==len(getColumn(y,0)) {
        for i,_ := range x{
            q := x[i]
            n := len(q)
            l := len(getColumn(x,i))
            for j := 0; j < n;j++{
                r := getColumn(y,j)
                
                p[i][j] =0
                for k:=0; k < l; k++{
                    p[i][j] =p[i][j] + q[k]*r[k] 
                }

            }      
        }
    }
    return p
}