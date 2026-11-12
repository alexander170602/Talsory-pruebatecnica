package models

// MatrixRequest representa la matriz recibida por la API.
type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

// QRResponse representa el resultado de la factorización QR.
type QRResponse struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}
