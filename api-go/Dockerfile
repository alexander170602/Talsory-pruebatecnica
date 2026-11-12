# Etapa de compilación
FROM golang:alpine AS builder

WORKDIR /app

# Copiamos los archivos de dependencias
COPY go.mod go.sum ./

RUN go mod download

# Copiamos el código fuente
COPY . .

# Compilamos la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Etapa de ejecución
FROM alpine:latest

WORKDIR /app

# Certificados para conexiones HTTP/HTTPS
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server .

ENV PORT=3001

EXPOSE 3001

CMD ["./server"]