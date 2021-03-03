package main

import (
	"fmt"
	"runtime/debug"
)

func r() {

}
func a() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("msg: ", r)
			debug.PrintStack()
		}
	}()

	n := []int{5, 7, 4, 5}
	if len(n) == 3 {
		fmt.Println("3")
	} else {
		//log.Fatal("x")
		panic(fmt.Errorf("error %s", "x"))
	}
	fmt.Println("normally returned from a")
}
func main() {
	a()
	fmt.Println("normally returned from main")
}
