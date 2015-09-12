// copyright 2015 hugo_gao  all rights reseved
// the Stack  test case

package  stack

import "testing"

import (
    "fmt"
    "strconv"
)
func TestStack(t *testing.T){
    s := new(Stack)
    s.Push(5)
    ret := s.Pop()
    if  ret !=5 {
        fmt.Printf("Pop doesn't gvie 5, it's %d \n", ret )
        t.Log("Pop doesn't gvie 5, it's  "+ strconv.Itoa( ret) )
        t.Fail()
    }
}

