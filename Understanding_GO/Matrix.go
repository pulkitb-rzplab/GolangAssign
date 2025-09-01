package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	rows int
	cols int
	ele  [][]int
}

func getR(m Matrix) int {
	return m.rows

}

func getC(m Matrix) int {
	return m.cols
}

func setE(m Matrix, i, j, val int) {
	m.ele[i][j] = val
	fmt.Println("Element at position", i, j, "set to", val)
	fmt.Println("Do you want to print the matrix?[Y/N]")
	a := ""
	fmt.Scanln(&a)
	switch a {
	case "Y", "y":
		fmt.Println(m)
	case "N", "n":
		fmt.Println("Matrix not printed")
		return
	default:
		fmt.Println("Invalid input")
	}
}

func addMat(m1, m2 Matrix) Matrix {
	//given that dimensions are correct
	sum := Matrix{rows: 3, cols: 3, ele: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}}
	for i := range m1.rows {
		for j := range m1.cols {
			print(i, " ", j, "\n")
			sum.ele[i][j] = m1.ele[i][j] + m2.ele[i][j]
		}

	}
	return sum
}

func printJson(m Matrix) {
	JSData, err := json.Marshal(m.ele)
	if err != nil {
		fmt.Println("Error during marshall process\n Exit")
		return
	}
	fmt.Println(string(JSData))
}

//func main() {
	matrix1 := Matrix{rows: 3, cols: 3, ele: [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}
	matrix2 := Matrix{rows: 3, cols: 3, ele: [][]int{
		{1, 2, 3},
		{8, 6, 7},
		{9, 11, 12},
	}}

	fmt.Println("Matrix 1:\n", matrix1.ele)
	fmt.Println("Matrix 2:\n", matrix2.ele)

	// fmt.Println("Rows in mat1:", getR(matrix1))
	// fmt.Println("columns in Mat1:", getC(matrix1))

	// fmt.Println("Setting element at position 1,1 to 15")
	// setE(matrix1, 1, 1, 15)
	// fmt.Println("Matrix 1 after setting element:", matrix1.ele)

	// var x, y, el int
	// fmt.Println("Enter the row and column where element need to be set")
	// fmt.Scanln(&x, &y)
	// fmt.Println("Enter the element to be set")
	// fmt.Scanln(&el)
	// setE(matrix1, x, y, el)
	// fmt.Println("Matrix 1 after setting element:", matrix1.ele)

	sumMat := addMat(matrix1, matrix2)
	fmt.Println("Sum of matrices:", sumMat.ele)
	fmt.Println()
	printJson(matrix1)
	printJson(matrix2)
	printJson(sumMat)
}
