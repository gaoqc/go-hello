package  fibonacci



func  Fibonacci(val int) []int{
     x := make([]int ,val)
     x[0], x[1] = 1, 1
     for n := 2; n < val ; n++ {
         x[n] = x[n-1] + x[n-2]
     }
     return x
 }

