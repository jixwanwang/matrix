package matrix

import "fmt"

type Vector struct {
	values []float64
}

func NewVector(size int) *Vector {
	return &Vector{
		values: make([]float64, size),
	}
}

func (V *Vector) Set(i int, v float64) {
	if len(V.values) > i {
		V.values[i] = v
	}
}

func (V *Vector) Get(i int) float64 {
	if len(V.values) > i {
		return V.values[i]
	}
	return 0.0
}

func (V *Vector) Print() {
	for _, v := range V.values {
		fmt.Printf("%v,", v)
	}
	fmt.Printf("\n")
}

type BooleanVector struct {
	values []bool
}

func NewBooleanVector(size int) *BooleanVector {
	return &BooleanVector{
		values: make([]bool, size),
	}
}

func (V *BooleanVector) Set(i int, v bool) {
	if len(V.values) > i {
		V.values[i] = v
	}
}

func (V *BooleanVector) Get(i int) bool {
	if len(V.values) > i {
		return V.values[i]
	}
	return false
}
