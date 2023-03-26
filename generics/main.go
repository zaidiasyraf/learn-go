package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// no generic addtions
	fmt.Printf("Non-Generic sums: %v and %v\n", SumInts(ints), SumFloats(floats))

	// Generic additions with type argument
	fmt.Printf("Generic sums with type argument: %v and %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))

	// Generic additions without type argument
	fmt.Printf("Generic sums with type inferred: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))

	// Generic sums with constraint
	fmt.Printf("Generic sums with constraint: %v and %v\n", SumNumbers(ints), SumNumbers(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
