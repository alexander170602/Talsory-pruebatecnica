
$ErrorActionPreference = "Stop"

Write-Host "Iniciando pruebas de integracion..." -ForegroundColor Cyan

$url = "http://localhost:3001/api/v1/qr"

$body = @{
    matrix = @(
        @(1, 2),
        @(3, 4),
        @(5, 6)
    )
} | ConvertTo-Json -Depth 5

try {
    $response = Invoke-RestMethod `
        -Uri $url `
        -Method Post `
        -ContentType "application/json" `
        -Body $body

    # Validar estado de la respuesta
    if ($response.status -ne "success") {
        throw "La API Go no devolvio status success"
    }

    $data = $response.data

    # Validar dimensiones de Q
    if ($data.q.Count -ne 3 -or $data.q[0].Count -ne 3) {
        throw "La matriz Q no tiene dimensiones 3x3"
    }

    # Validar dimensiones de R
    if ($data.r.Count -ne 3 -or $data.r[0].Count -ne 2) {
        throw "La matriz R no tiene dimensiones 3x2"
    }

    # Validar que Node.js haya devuelto las estadisticas
    $statistics = $data.statistics

    $requiredFields = @(
        "max",
        "min",
        "average",
        "sum",
        "qIsDiagonal",
        "rIsDiagonal",
        "isDiagonal"
    )

    foreach ($field in $requiredFields) {
        if ($null -eq $statistics.$field) {
            throw "Falta el campo de estadistica: $field"
        }
    }

    Write-Host ""
    Write-Host "TODAS LAS PRUEBAS DE INTEGRACION PASARON" -ForegroundColor Green
    Write-Host "API Go: OK"
    Write-Host "API Node.js: OK"
    Write-Host "Comunicacion HTTP: OK"
    Write-Host "Factorizacion QR: OK"
    Write-Host "Estadisticas: OK"

    Write-Host ""
    Write-Host "Resultado de estadisticas:"
    $statistics | Format-List

}
catch {
    Write-Host ""
    Write-Host "ERROR EN LA PRUEBA DE INTEGRACION" -ForegroundColor Red
    Write-Host $_.Exception.Message -ForegroundColor Red
    exit 1
}
