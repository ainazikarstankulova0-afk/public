package main

import "fmt"

func main() {
	x:=66
	inc(&x)
	fmt.Println(x)
}

func inc(x *int){
	*x++
}