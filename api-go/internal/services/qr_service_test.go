package services

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestFactorizeQR(t *testing.T) {
	input := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	q, r, err := FactorizeQR(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Convert the original matrix into Gonum format.
	a := mat.NewDense(3, 2, []float64{
		1, 2,
		3, 4,
		5, 6,
	})

	// Convert Q and R into Gonum matrices.
	qMatrix := mat.NewDense(3, 3, flatten(q))
	rMatrix := mat.NewDense(3, 2, flatten(r))

	// Verify that Q * R reconstructs A.
	var reconstructed mat.Dense
	reconstructed.Mul(qMatrix, rMatrix)

	if !mat.EqualApprox(a, &reconstructed, 1e-10) {
		t.Error("Q * R does not reconstruct the original matrix")
	}

	// Verify that Q transpose * Q equals the identity matrix.
	var orthogonality mat.Dense
	orthogonality.Mul(qMatrix.T(), qMatrix)

	identity := mat.NewDense(3, 3, []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	})

	if !mat.EqualApprox(&orthogonality, identity, 1e-10) {
		t.Error("Q is not orthogonal")
	}
}

func TestFactorizeQRRejectsInvalidMatrices(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]float64
	}{
		{
			name:   "empty matrix",
			matrix: [][]float64{},
		},
		{
			name:   "empty row",
			matrix: [][]float64{{}},
		},
		{
			name: "irregular matrix",
			matrix: [][]float64{
				{1, 2},
				{3},
			},
		},
		{
			name: "wide matrix",
			matrix: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, _, err := FactorizeQR(test.matrix)

			if err == nil {
				t.Error("expected an error for invalid matrix")
			}
		})
	}
}

func flatten(matrix [][]float64) []float64 {
	var result []float64

	for _, row := range matrix {
		result = append(result, row...)
	}

	return result
}
