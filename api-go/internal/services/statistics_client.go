package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// StatisticsRequest representa los datos que enviaremos a Node.js.
type StatisticsRequest struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}

// StatisticsResult representa las estadísticas recibidas de Node.js.
type StatisticsResult struct {
	Max         float64 `json:"max"`
	Min         float64 `json:"min"`
	Average     float64 `json:"average"`
	Sum         float64 `json:"sum"`
	QIsDiagonal bool    `json:"qIsDiagonal"`
	RIsDiagonal bool    `json:"rIsDiagonal"`
	IsDiagonal  bool    `json:"isDiagonal"`
}

// GetStatistics envía Q y R a la API Node.js.
func GetStatistics(q, r [][]float64) (*StatisticsResult, error) {

	// URL configurable mediante variable de entorno.
	apiURL := os.Getenv("STATISTICS_API_URL")

	if apiURL == "" {
		apiURL = "http://localhost:3002/api/v1/statistics"
	}

	// Preparar los datos JSON.
	requestData := StatisticsRequest{
		Q: q,
		R: r,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error encoding statistics request: %w", err)
	}

	// Crear una petición con tiempo límite.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		apiURL,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating statistics request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Ejecutar petición HTTP.
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error connecting to Node.js API: %w", err)
	}
	defer resp.Body.Close()

	// Comprobar respuesta HTTP.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"Node.js API returned HTTP status %d",
			resp.StatusCode,
		)
	}

	// Estructura de la respuesta de Node.js.
	var response struct {
		Status string           `json:"status"`
		Data   StatisticsResult `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding statistics response: %w", err)
	}

	if response.Status != "success" {
		return nil, fmt.Errorf("Node.js API returned an unsuccessful response")
	}

	return &response.Data, nil
}
