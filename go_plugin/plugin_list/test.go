package main

import "fmt"

func Run(params []interface{}) {
	fmt.Println(params)
}

func main()  {
	//go build -buildmode=c-shared -o test.so
	fmt.Println("test plugin")
}
