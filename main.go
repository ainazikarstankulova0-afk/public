package main

import "fmt"

func main() {
	x:=5
	Ukazatel(&x)
	fmt.Println(x)
}
 
func Ukazatel(*x int){
	*x=*x+1
}