package main

import(
    "fmt"
//    "github.com/gaoqc/stringutil"
//    "github.com/gaoqc/fibonacci"
)
type T struct{
 name  string
 age int
 boy bool
}
func main(){
    //fmt.Printf(stringutil.Reverse("hello,go!\n"))

   // fmt.Printf("%v\n",fibonacci.Fibonacci(10))
    // f := plusTwo()
    //fmt.Printf("%d\n", plusTwo()(2))
  for  d,_ := range make([]int,100){
      fmt.Printf("%d,",d)
  }
}
func  plusTwo() func(int) int{
    return  func(n int) int{return n+2}
}
