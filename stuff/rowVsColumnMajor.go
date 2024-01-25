package stuff

import (
	"fmt"
	"strconv"
	"time"
)

type Foo struct {
	X int
	F float32
}

/*
Row major is the one that makes smaller jumps; 
Column major during each iteration indexes addresses that are size apart
*/
func SumRowMajor(matrix []int, size int) int64 {
	defer Timer("Row Major, size " + strconv.Itoa(size))()
	var sum int64 = 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += int64(matrix[i*size+j])
		}
	}
	return sum
}

func SumFooRowMajor(matrix []Foo, size int) int64 {
	defer Timer("Foo Row Major, size " + strconv.Itoa(size))()
	var sum int64 = 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += int64(matrix[i*size+j].X)
		}
	}
	return sum
}

func SumFooNestedRowMajor(matrix [][]Foo, size int) int64 {
	defer Timer("Foo Nested Row Major, size " + strconv.Itoa(size))()
	var sum int64 = 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += int64(matrix[i][j].X)
		}
	}
	return sum
}

func SumColumnMajor(matrix []int, size int) int64 {
	defer Timer("Col Major, size " + strconv.Itoa(size))()
	var sum int64 = 0
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			sum += int64(matrix[i*size+j])
		}
	}
	return sum
}

func SumColumnMajorNested(matrix [][]int, size int) int64 {
	defer Timer("Nested Col Major, size " + strconv.Itoa(size))()
	var sum int64 = 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += int64(matrix[j][i])
		}
	}
	return sum
}

func SumRowMajorNested(matrix [][]int, size int) int64 {
	defer Timer("Nested Row Major, size " + strconv.Itoa(size))()
	var sum int64 = 0
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			sum += int64(matrix[j][i])
		}
	}
	return sum
}


func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v micros\n", name, time.Since(start).Microseconds())
	}
}
