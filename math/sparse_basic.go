package math

import "math"

/*
Swap two rows in this matrix.
*/
func (A *SparseMatrix) SwapRows(r1, r2 uint) {
	js := map[uint]bool{}
	for index := range A.elements {
		i, j := A.GetRowColIndex(index)
		if i == r1 || i == r2 {
			js[j] = true
		}
	}
	for j := range js {
		tmp := A.Get(r1, j)
		A.Set(r1, j, A.Get(r2, j))
		A.Set(r2, j, tmp)
	}
}

/*
Scale a row by a scalar.
*/
func (A *SparseMatrix) ScaleRow(r uint, f float64) {
	for index, value := range A.elements {
		i, j := A.GetRowColIndex(index)
		if i == r {
			A.Set(i, j, value*f)
		}
	}
}

/*
Add a multiple of row rs to row rd.
*/
func (A *SparseMatrix) ScaleAddRow(rd, rs uint, f float64) {
	//	for index, value := range A.elements {
	//		i, j := A.GetRowColIndex(index)
	//		if i == rs {
	//			A.Set(rd, j, A.Get(rd, j)+value*f)
	//		}
	//	}

	for j := uint(0); j < A.cols; j++ {
		if val, ok := A.Get(rs, j); ok {
			if old, okd := A.Get(rdd, j); okd {
				val += old
			}
			A.Set(rd, j, val)
		}
	}
}

func (A *SparseMatrix) Symmetric() bool {
	for index, value := range A.elements {
		if i, j := A.GetRowColIndex(index); i != j {
			if val, ok := A.Get(j, i); ok != nil || value != val {
				return false
			}
		}
	}
	return true
}

func (A *SparseMatrix) Transpose() *SparseMatrix {
	B := ZerosSparse(A.cols, A.rows)
	for index, value := range A.elements {
		i, j := A.GetRowColIndex(index)
		B.Set(j, i, value)
	}
	return B
}

func (A *SparseMatrix) Det() float64 {
	//TODO: obviously this is a horrible way to do it
	return A.DenseMatrix().Det()
}

func (A *SparseMatrix) Trace() (res float64) {
	for index, value := range A.elements {
		i, j := A.GetRowColIndex(index)
		if i == j {
			res += value
		}
	}
	return
}

func (A *SparseMatrix) OneNorm() (res float64) {
	for _, value := range A.elements {
		res += math.Abs(value)
	}
	return
}

func (A *SparseMatrix) TwoNorm() float64 {
	var sum float64 = 0
	for _, value := range A.elements {
		sum += value * value
	}
	return math.Sqrt(sum)
}

func (A *SparseMatrix) InfinityNorm() (res float64) {
	for _, value := range A.elements {
		res = max(res, math.Abs(value))
	}
	return
}
