// copyright 2015  hugo_gao  all rights  reserved
// package stack  provides  FILO  and  Push and Pop method
package  stack

// define Stack type
type  Stack struct{
     index int
     array [10]int
}
// Push  data to  stack
func (s *Stack) Push( num int ) bool{
    if  s.index >= 10 {
        return false
    }
    s.array[s.index] =  num
    s.index ++
    return true
}
// pop  data of last push
func (s *Stack) Pop()( ret int){
    if s.index<0 {
        return -1
    }
    s.index --
    ret =s.array[s.index]
    return
}

