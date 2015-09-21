package main

import "fmt"

func main() {
	mat := [][]int32{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	rows := len(mat)
	cols := len(mat[0])
	fmt.Printf("# original: %d x %d\n", rows, cols)
	fmt.Println(mat)

	trrows := cols
	trcols := rows
	trmat := make([][]int32, trrows)
	for trr := 0; trr < trrows; trr++ {
		trmat[trr] = make([]int32, trcols)
	}

	for trr := 0; trr < trrows; trr++ {
		for trc := 0; trc < trcols; trc++ {
			trmat[trr][trc] = mat[trc][trr]
		}
	}

	fmt.Printf("# transposed: %d x %d\n", len(trmat), len(trmat[0]))
	fmt.Println(trmat)
}
