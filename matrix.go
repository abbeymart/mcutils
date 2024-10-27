package mcutils

import (
	"errors"
	"fmt"
)

// AddMatrices function adds two matrices of the same dimensions.
func AddMatrices[T float64 | int64](matrix1 [][]T, matrix2 [][]T, result [][]T) error {
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
		mat1 := matrix1[matrixIndex]
		mat2 := matrix2[matrixIndex]
		// validate matrix1 and matrix2 sub-items length
		if len(mat1) != subItemLength || len(mat2) != subItemLength {
			result = [][]T{}
			return errors.New(fmt.Sprintf("length of both sub-matrices should be equal [matrix1[%v]: %v | matrix2[%v]: %v]", matrix1[matrixIndex], len(matrix1), matrix2[matrixIndex], len(matrix2)))
		}
		// compute matrix additions
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

// AddMultipleMatrices function [tensor] adds multiple matrices of the same dimensions.
func AddMultipleMatrices[T float64 | int64](matrices [][][]T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	matricesLength := len(matrices)
	if matricesLength <= 1 {
		return errors.New(fmt.Sprintf("length of matrices should be greater than 1"))
	}
	// perform addition of the first two matrices
	err := AddMatrices(matrices[0], matrices[1], result)
	if err != nil {
		result = [][]T{}
		return err
	}
	// perform the remaining addition of the 3rd to the last matrix
	matIndex := 2
	for matIndex < matricesLength {
		var nextResult [][]T
		err = AddMatrices(result, matrices[matIndex], nextResult)
		if err != nil {
			result = [][]T{}
			return err
		}
		result = nextResult
		matIndex += 1
	}
	return nil
}

// SubtractMatrices function subtract two matrices of the same dimensions.
func SubtractMatrices[T float64 | int64](matrix1 [][]T, matrix2 [][]T, result [][]T) error {
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
		mat1 := matrix1[matrixIndex]
		mat2 := matrix2[matrixIndex]
		// validate matrix1 and matrix2 sub-items length
		if len(mat1) != subItemLength || len(mat2) != subItemLength {
			result = [][]T{}
			return errors.New(fmt.Sprintf("length of both sub-matrices should be equal [matrix1[%v]: %v | matrix2[%v]: %v]", matrix1[matrixIndex], len(matrix1), matrix2[matrixIndex], len(matrix2)))
		}
		// compute matrix subtractions
		var matAddResult []T
		subItemIndex := 0
		for subItemIndex < subItemLength {
			// perform subtraction
			matAddResult = append(matAddResult, mat1[subItemIndex]-mat2[subItemIndex])
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

// SubtractMultipleMatrices function [tensor] subtract multiple matrices of the same dimensions.
func SubtractMultipleMatrices[T float64 | int64](matrices [][][]T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	matricesLength := len(matrices)
	if matricesLength <= 1 {
		return errors.New(fmt.Sprintf("length of matrices should be greater than 1"))
	}
	// perform subtraction of the first two matrices
	err := SubtractMatrices(matrices[0], matrices[1], result)
	if err != nil {
		result = [][]T{}
		return err
	}
	// perform the remaining subtraction of the 3rd to the last matrix
	matIndex := 2
	for matIndex < matricesLength {
		var nextResult [][]T
		err = SubtractMatrices(result, matrices[matIndex], nextResult)
		if err != nil {
			result = [][]T{}
			return err
		}
		result = nextResult
		matIndex += 1
	}
	return nil
}

// AddScalarMatrix function adds a scalar Value to the matrix/matrices.
func AddScalarMatrix[T float64 | int64](matrix [][]T, scalar T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	// validate matrix1 and matrix2 length
	if len(matrix) < 1 {
		return errors.New(fmt.Sprintf("length of the matrix should greater than 0"))
	}
	matrixLength := len(matrix)
	subItemLength := len(matrix[0])
	matrixIndex := 0
	for matrixIndex < matrixLength {
		mat := matrix[matrixIndex]
		// compute matrix additions
		var matAddResult []T
		subItemIndex := 0
		for subItemIndex < subItemLength {
			// perform addition
			matAddResult = append(matAddResult, mat[subItemIndex]+scalar)
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

// SubtractScalarMatrix function subtracts a scalar Value from the matrix/matrices.
func SubtractScalarMatrix[T float64 | int64](matrix [][]T, scalar T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	// validate matrix1 and matrix2 length
	if len(matrix) < 1 {
		return errors.New(fmt.Sprintf("length of the matrix should greater than 0"))
	}
	matrixLength := len(matrix)
	subItemLength := len(matrix[0])
	matrixIndex := 0
	for matrixIndex < matrixLength {
		mat := matrix[matrixIndex]
		// compute matrix additions
		var matAddResult []T
		subItemIndex := 0
		for subItemIndex < subItemLength {
			// perform addition
			matAddResult = append(matAddResult, mat[subItemIndex]-scalar)
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

// MultiplyScalarMatrix function multiply a scalar Value with the matrix/matrices.
func MultiplyScalarMatrix[T float64 | int64](matrix [][]T, scalar T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	// validate matrix1 and matrix2 length
	if len(matrix) < 1 {
		return errors.New(fmt.Sprintf("length of the matrix should greater than 0"))
	}
	matrixLength := len(matrix)
	subItemLength := len(matrix[0])
	matrixIndex := 0
	for matrixIndex < matrixLength {
		mat := matrix[matrixIndex]
		// compute matrix additions
		var matAddResult []T
		subItemIndex := 0
		for subItemIndex < subItemLength {
			// perform addition
			matAddResult = append(matAddResult, mat[subItemIndex]*scalar)
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

// DivideScalarMatrix function the matrix/matrices by the scalar Value.
func DivideScalarMatrix[T float64 | int64](matrix [][]T, scalar T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	// validate matrix1 and matrix2 length
	if len(matrix) < 1 {
		return errors.New(fmt.Sprintf("length of the matrix should greater than 0"))
	}
	matrixLength := len(matrix)
	subItemLength := len(matrix[0])
	matrixIndex := 0
	for matrixIndex < matrixLength {
		mat := matrix[matrixIndex]
		// compute matrix additions
		var matAddResult []T
		subItemIndex := 0
		for subItemIndex < subItemLength {
			// perform addition
			matAddResult = append(matAddResult, mat[subItemIndex]/scalar)
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

// TransposeMatrix function transpose the matrix - swap rows and columns, i.e. rotate the matrix around it's diagonal.
func TransposeMatrix[T float64 | int64](matrix [][]T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	// validate matrix length
	if len(matrix) < 1 {
		return errors.New(fmt.Sprintf("length of the matrix should greater than 0"))
	}
	for _, matSlice := range matrix {
		if len(matrix[0]) != len(matSlice) {
			return errors.New(fmt.Sprintf("Length of matrix2 sub-items must be equal [Expected: %v, Got: %v]", len(matrix[0]), len(matSlice)))
		}
	}
	// transpose matrix, swap columns to rows, diagonally
	matColumnItemsCount := len(matrix[0])
	matColumnItemIndex := 0
	for matColumnItemIndex < matColumnItemsCount {
		var transposeSliceRow []T
		for _, matColumnSlice := range matrix {
			transposeSliceRow = append(transposeSliceRow, matColumnSlice[matColumnItemIndex])
		}
		result = append(result, transposeSliceRow)
		matColumnItemIndex += 1
	}
	return nil
}

// MultiplyMatrix function multiply two matrices.
// The matrix1 single slice length must be the same as the number of columns in matrix2/sub-matrices.
func MultiplyMatrix[T float64 | int64](matrix1 []T, matrix2 [][]T, result []T) error {
	// initialize the matrix result
	result = []T{}
	// validate matrix2 values' lengths must match the length of matrix1[0]
	if len(matrix1) != len(matrix2) {
		return errors.New(fmt.Sprintf("Length of matrix1 [Expected: %v] must match the number of columns of matrix2 [Got: %v]", len(matrix1), len(matrix2)))
	}
	for _, mat2Slice := range matrix2 {
		if len(matrix2[0]) != len(mat2Slice) {
			return errors.New(fmt.Sprintf("Length of matrix2 sub-items must be equal [Expected: %v, Got: %v]", len(matrix2[0]), len(mat2Slice)))
		}
	}
	// compute the matrices multiplication
	mat1Slice := matrix1
	mat1Columns := len(mat1Slice) // ==> matrix2 sub-items length/columns
	mat1ColCount := 0
	var matMultiSlices [][]T // Required to compute the summation of the row-column multiplications
	for mat1ColCount < mat1Columns {
		// compose multiplication Slice, by matching matrix1/matrix2-columns
		mat1ColVal := mat1Slice[mat1ColCount]
		mat2ColSlice := matrix2[mat1ColCount]
		var matMultiSlice []T
		for _, mat2ColVal := range mat2ColSlice {
			matMultiSlice = append(matMultiSlice, mat2ColVal*mat1ColVal)
		}
		// update mat-multiplication-slice
		matMultiSlices = append(matMultiSlices, matMultiSlice)
		// next column
		mat1ColCount += 1
	}
	// compute the sum of multiplication-slices by matching columns/rows
	//var sumMultiplication []T
	matMultiRows := len(matMultiSlices[0])
	matMultiRow := 0
	for matMultiRow < matMultiRows {
		matMultiSum := T(0)
		for _, val := range matMultiSlices {
			matMultiSum += val[matMultiRow]
		}
		result = append(result, matMultiSum)
		// next row
		matMultiRow += 1
	}
	return nil
}

// MultiplyMatrices function multiply two matrices.
// The length of each of the matrix1 sub-matrix/rows must match the length of matrix2 (columns).
func MultiplyMatrices[T float64 | int64](matrix1 [][]T, matrix2 [][]T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	// validate matrix1 sub-items and matrix2 length, rows/columns matching
	for _, matrix1Val := range matrix1 {
		if len(matrix1[0]) != len(matrix1Val) {
			return errors.New(fmt.Sprintf("Length of matrix1 sub-items must be the same [Expected: %v, Got: %v]", len(matrix1[0]), len(matrix1Val)))
		}
		if len(matrix1Val) != len(matrix2) {
			return errors.New(fmt.Sprintf("Length of matrix1 sub-items must match the matrix2 columns/length [Expected: %v, Got: %v]", len(matrix1Val), len(matrix2)))
		}
	}
	// validate matrix2 sub-items lengths/rows
	for _, mat2Slice := range matrix2 {
		if len(matrix2[0]) != len(mat2Slice) {
			return errors.New(fmt.Sprintf("Length of matrix2 sub-items must be equal [Expected: %v, Got: %v]", len(matrix2[0]), len(mat2Slice)))
		}
	}
	// compute the matrices multiplication
	matrix1SlicesLength := len(matrix1)
	matrix1SliceIndex := 0
	for matrix1SliceIndex < matrix1SlicesLength {
		// compute the matrix multiplication for each of the matrix1 slices and matrix2 slices
		var multiResult []T
		err := MultiplyMatrix(matrix1[matrix1SliceIndex], matrix2, multiResult)
		if err != nil {
			result = [][]T{}
			return err
		}
		result = append(result, multiResult)
		matrix1SliceIndex += 1
	}
	return nil
}
