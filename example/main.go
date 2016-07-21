package main

import "github.com/jixwanwang/matrix"

func main() {
	m := matrix.NewSquareSparseMatrix(4)
	m.Set(1, 0, 5)
	m.Set(1, 1, 8)
	m.Set(2, 2, 3)
	m.Set(3, 1, 6)
	m.Print()
	m.Compress()
	m.Print()

	v := matrix.NewBooleanVector(4)
	v.Set(1, true)
	v.Set(2, true)
	v.Set(3, true)

	vv := m.Multiply(v)
	vv.Print()
}
