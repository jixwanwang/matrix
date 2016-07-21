package matrix

import (
	"fmt"
	"sort"
)

type SparseMatrix interface {
	Set(r, c int, v float64)
	Get(r, c int) float64
	Compress()
	Multiply(v *BooleanVector) *Vector
}

type element struct {
	r int
	c int
	v float64
}

type elements []*element

func (E elements) Len() int {
	return len(E)
}

func (E elements) Swap(i, j int) {
	E[i], E[j] = E[j], E[i]
}

func (E elements) Less(i, j int) bool {
	if E[i].r != E[j].r {
		return E[i].r < E[j].r
	} else {
		return E[i].c < E[j].c
	}
	return false
}

type SquareCSRMatrix struct {
	data      []float64
	columns   []int
	rowStarts []int

	size int

	tempElements elements
	compressed   bool
}

func NewSquareSparseMatrix(size int) *SquareCSRMatrix {
	return &SquareCSRMatrix{
		size:         size,
		tempElements: elements([]*element{}),
	}
}

func (M *SquareCSRMatrix) Set(r, c int, v float64) {
	M.tempElements = append(M.tempElements, &element{r: r, c: c, v: v})
}

// Compress calculates the CSR representation of a sparse matrix.
func (M *SquareCSRMatrix) Compress() {
	// Put elements in row major order
	sort.Sort(M.tempElements)

	// Make arrays
	M.data = make([]float64, len(M.tempElements))
	M.columns = make([]int, len(M.tempElements))
	M.rowStarts = make([]int, M.size+1)

	// Set first and last element of row starts
	M.rowStarts[0] = 0
	M.rowStarts[M.size] = len(M.tempElements)

	// Row we're currently going through
	currRow := 0
	for i, e := range M.tempElements {
		M.data[i] = e.v
		M.columns[i] = e.c
		if e.r > currRow {
			// The number of elements seen up to this row is i
			// Set all the rowStarts since the row we were most recently processing to this.
			for r := 1; r < e.r-currRow+1; r++ {
				M.rowStarts[currRow+r] = i
			}
			currRow = e.r
		}
	}

	M.compressed = true
}

func (M *SquareCSRMatrix) Multiply(v *BooleanVector) *Vector {
	if len(v.values) != M.size {
		return nil
	}

	// Only care about rows that correspond to a true in the vector
	// output := NewVector(M.size)
	output := make([]float64, M.size)
	rows := []int{}
	for i := 0; i < M.size; i++ {
		if v.Get(i) {
			rows = append(rows, i)
		}
	}

	for _, r := range rows {
		for d := M.rowStarts[r]; d < M.rowStarts[r+1]; d++ {
			c := M.columns[d]
			output[c] += M.data[d]
		}
	}

	vector := NewVector(M.size)
	for i, v := range output {
		vector.Set(i, v)
	}

	return vector
}

func (M *SquareCSRMatrix) Print() {
	if M.compressed {
		fmt.Printf("Data: %v\n", M.data)
		fmt.Printf("Rows: %v\n", M.rowStarts)
		fmt.Printf("Columns: %v\n", M.columns)
	} else {
		sort.Sort(M.tempElements)
		for _, e := range M.tempElements {
			fmt.Printf("%d %d %v\n", e.r, e.c, e.v)
		}
	}
}
