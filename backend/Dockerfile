# Etapa 1: Compilar el binario con Go
FROM golang:1.24.3-alpine AS builder

# Instalar dependencias del sistema necesarias
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar el binario con optimizaciones
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Etapa 2: Crear imagen liviana para producción
FROM alpine:latest

# Instalar dependencias necesarias
RUN apk --no-cache add ca-certificates tzdata wget curl

WORKDIR /app

# Copiar binario desde el builder
COPY --from=builder /app/app .

# Crear usuario no-root para seguridad
RUN adduser -D -s /bin/sh appuser && chown -R appuser:appuser /app
USER appuser

# Exponer el puerto en el que escucha el backend (Gin)
EXPOSE 8080

# Comando para iniciar la app
CMD ["./app"]
