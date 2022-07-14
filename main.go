package main

import (
	"fmt"
)

type Iterator[T any] interface {
	Each(fn func(value T, index int))
}

type Transformer[T any, R any] interface {
	Map(fn func(value T, index int) R) []R
}

type Aggregator[T any, R any] interface {
	Reduce(fn func(accumulator R, value T, index int) R) R
}

type Iterable[T any, R any] interface {
	Iterator[T]
	Transformer[T, R]
	Aggregator[T, R]
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	fmt.Println("hellow orld")
	r := makeRange(5, 25)
	for i, k := range r {
		fmt.Printf("number[%d]=%d\n", i, k)
	}
}
