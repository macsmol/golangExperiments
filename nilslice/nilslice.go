package main

import "fmt"


type foo struct {
	a int
	someSlice []int
}

func main() {
	var slice []int
	fmt.Println("len", len(slice))

	var myStruct foo
	fmt.Println("len from slice in zero struct ", len(myStruct.someSlice))

	var myStructPtr *foo = nil
	fmt.Println("len from slice in nil struct ptr", len(myStructPtr.someSlice))

}