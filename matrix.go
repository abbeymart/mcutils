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
			return errors.New(fmt.Sprintf("length of both sub-matrices should be equal [matrix1[%v]: %v | matrix2[%v]: %v]", matrixIndex, len(matrix1), matrixIndex, len(matrix2)))
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

// AddMultipleMatrices function adds multiple matrices of the same dimensions.
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
			return errors.New(fmt.Sprintf("length of both sub-matrices should be equal [matrix1[%v]: %v | matrix2[%v]: %v]", matrixIndex, len(matrix1), matrixIndex, len(matrix2)))
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

// SubtractMultipleMatrices function subtract multiple matrices of the same dimensions.
func SubtractMultipleMatrices[T float64 | int64](matrices [][][]T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	matricesLength := len(matrices)
	if matricesLength <= 1 {
		return errors.New(fmt.Sprintf("length of matrices should be greater than 1"))
	}
	// perform subtraction of the first two matrices
	err := AddMatrices(matrices[0], matrices[1], result)
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

// AddScalarMatrix function adds a scalar value to the matrix/matrices.
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

// SubtractScalarMatrix function subtracts a scalar value from the matrix/matrices.
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

// MultiplyScalarMatrix function multiply a scalar value with the matrix/matrices.
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

// DivideScalarMatrix function the matrix/matrices by the scalar value.
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
// TODO: complete
func TransposeMatrix[T float64 | int64](matrix [][]T, result [][]T) error {
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

// MultiplyMatrices function multiply two matrices. TODO: complete
// The number of rows in matrix1 must be the same as the number of columns in matrix2.
func MultiplyMatrices[T float64 | int64](matrix1 [][]T, matrix2 [][]T, result [][]T) error {
	// initialize the matrix result
	result = [][]T{}
	// validate matrix1 and matrix2 length
	matrixLength := len(matrix2)
	if len(matrix1) > 1 || len(matrix1[0]) != len(matrix2[0]) {
		return errors.New(fmt.Sprintf("length of matrix should be 1. Length of matrix1[0] must be the same as the number of columns in matrix2 [%v]", len(matrix2[0])))
	}
	subItemLength := len(matrix1[0])
	matrixIndex := 0
	for matrixIndex < matrixLength {
		mat1 := matrix1[matrixIndex]
		mat2 := matrix2[matrixIndex]
		// validate matrix1 and matrix2 sub-items length
		if len(mat1) != subItemLength || len(mat2) != subItemLength {
			result = [][]T{}
			return errors.New(fmt.Sprintf("length of both sub-matrices should be equal [matrix1[%v]: %v | matrix2[%v]: %v]", matrixIndex, len(matrix1), matrixIndex, len(matrix2)))
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
