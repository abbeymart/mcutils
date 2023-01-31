package mcutils

import (
	"errors"
	"fmt"
)

// AddMatrix function adds two matrices of the same dimensions.
func AddMatrix[T float64 | int64](matrix1 [][]T, matrix2 [][]T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	// validate matrix1 and matrix2 length
	if len(matrix1) != len(matrix2) {
		return errors.New(fmt.Sprintf("length of both matrices should be equal [matrix1: %v | matrix2: %v]", len(matrix1), len(matrix2)))
	}
	matrixLength := len(matrix1)
	subItemLength := len(matrix1[0])
	matrixIndex := 0
	for matrixIndex < matrixLength {
		// validate matrix1 and matrix2 sub-items length
		if len(matrix1[matrixIndex]) != subItemLength || len(matrix2[matrixIndex]) != subItemLength {
			result = [][]T{}
			return errors.New(fmt.Sprintf("length of both sub-matrices should be equal [matrix1[%v]: %v | matrix2[%v]: %v]", matrixIndex, len(matrix1), matrixIndex, len(matrix2)))
		}
		// compute matrix additions
		mat1 := matrix1[matrixIndex]
		mat2 := matrix2[matrixIndex]
		var matAddResult []T
		subItemIndex := 0
		for subItemIndex < subItemLength {
			// perform addition
			matAddResult = append(matAddResult, mat1[subItemIndex]+mat2[subItemIndex])
			// increment subItemIndex
			subItemIndex += 1
		}
		// update result
		result = append(result, matAddResult)
		// increment matrixIndex
		matrixIndex += 1
	}
	return nil
}

// AddMultipleMatrix function adds multiple matrices of the same dimensions.
func AddMultipleMatrix[T float64 | int64](matrices [][][]T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	matricesLength := len(matrices)
	if matricesLength <= 1 {
		return errors.New(fmt.Sprintf("length of matrices should be greater than 1"))
	}
	// perform addition of the first two matrices
	err := AddMatrix(matrices[0], matrices[1], result)
	if err != nil {
		result = [][]T{}
		return err
	}
	// perform the remaining addition of the 3rd to the last matrix
	matIndex := 2
	for matIndex < matricesLength {
		var nextResult [][]T
		err = AddMatrix(result, matrices[matIndex], nextResult)
		if err != nil {
			result = [][]T{}
			return err
		}
		result = nextResult
		matIndex += 1
	}
	return nil
}
