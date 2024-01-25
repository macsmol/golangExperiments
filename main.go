package main

import (
	"fmt"
	"macsmol/quickCheckSth/stuff"
	"math/rand"
	"strconv"
)

var (
	SizeSerie = [...]int{
		10   , 20   , 50  ,
		100  , 200  , 500 ,
		1000 , 2000 , 5000,
		10000, 20000}
)

func main() {
	for _, size := range SizeSerie {
		matrix := generateTestData(size)
		
		stuff.SumRowMajor(matrix, size)
		stuff.SumColumnMajor(matrix, size)

		// transform []int  ->  [][]int
		var matrixNested [][]int = make([][]int, size)
		for i := range matrixNested {
			matrixNested[i], matrix = matrix[:size], matrix[size:]
		}

		stuff.SumRowMajorNested(matrixNested, size)
		stuff.SumColumnMajorNested(matrixNested, size)

		fmt.Println()
	}
	
	fmt.Println("-----------------------")

	for _, size := range SizeSerie {
		matrix := generateTestFoos(size)
		fmt.Println("foo ", stuff.SumFooRowMajor(matrix, size))

		
		var matrixNested [][]stuff.Foo = make([][]stuff.Foo, size)
		for i := range matrixNested {
			matrixNested[i], matrix = matrix[:size], matrix[size:]
		}
		fmt.Println("foo nested", stuff.SumFooNestedRowMajor(matrixNested, size))

		fmt.Println()
	}
}

func generateTestData(size int) []int {
	defer stuff.Timer("generateTestData(" + strconv.Itoa(size) + ")")()
	matrix := make([]int, 0, size*size)
	for i := 0; i < cap(matrix); i++ {
		matrix = append(matrix, rand.Intn(1000))
	}
	
	return matrix
}

func generateTestFoos(size int) []stuff.Foo {
	defer stuff.Timer("generateTestFoos(" + strconv.Itoa(size) + ")")()
	matrix := make([]stuff.Foo, 0, size*size)
	for i := 0; i < cap(matrix); i++ {
		matrix = append(matrix, stuff.Foo{rand.Intn(1000), 1.0})
	}
	
	return matrix
}
