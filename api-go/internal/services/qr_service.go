package services

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// FactorizeQR calcula la factorización QR de una matriz.
//
// Utilizamos Gonum para realizar la descomposición
// mediante el método de Householder.
func FactorizeQR(matrix [][]float64) ([][]float64, [][]float64, error) {

	// Validar que la matriz tenga filas.
	if len(matrix) == 0 {
		return nil, nil, fmt.Errorf("matrix cannot be empty")
	}

	rows := len(matrix)
	cols := len(matrix[0])

	// Validar que existan columnas.
	if cols == 0 {
		return nil, nil, fmt.Errorf("matrix cannot have empty rows")
	}

	// Gonum QR requiere filas >= columnas.
	if rows < cols {
		return nil, nil, fmt.Errorf(
			"matrix must have at least as many rows as columns",
		)
	}

	// Convertir la matriz bidimensional a un arreglo lineal.
	data := make([]float64, 0, rows*cols)

	for i, row := range matrix {

		// Validar que todas las filas tengan igual cantidad de columnas.
		if len(row) != cols {
			return nil, nil, fmt.Errorf(
				"matrix must be rectangular: row %d has an invalid size",
				i,
			)
		}

		for _, value := range row {

			// Rechazar valores infinitos o NaN.
			if math.IsNaN(value) || math.IsInf(value, 0) {
				return nil, nil, fmt.Errorf(
					"matrix contains an invalid numeric value",
				)
			}

			data = append(data, value)
		}
	}

	// Crear matriz densa de Gonum.
	a := mat.NewDense(rows, cols, data)

	// Ejecutar factorización QR.
	var qr mat.QR
	qr.Factorize(a)

	// Extraer matriz ortogonal Q.
	var q mat.Dense
	qr.QTo(&q)

	// Extraer matriz triangular superior R.
	var r mat.Dense
	qr.RTo(&r)

	// Convertir Q y R a estructuras JSON compatibles.
	qMatrix := denseToSlice(&q)
	rMatrix := denseToSlice(&r)

	return qMatrix, rMatrix, nil
}

// denseToSlice convierte una matriz Gonum a [][]float64.
func denseToSlice(matrix *mat.Dense) [][]float64 {

	rows, cols := matrix.Dims()

	result := make([][]float64, rows)

	for i := 0; i < rows; i++ {

		result[i] = make([]float64, cols)

		for j := 0; j < cols; j++ {
			result[i][j] = matrix.At(i, j)
		}
	}

	return result
}
