# Utiliza una imagen de Go para compilar
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum a la imagen
COPY go.mod go.sum ./

# Descarga las dependencias del módulo Go
RUN go mod download

# Copia el código fuente a la imagen
COPY src/ ./src/

# Compila el código en la carpeta build
RUN mkdir /app/build
CMD go build -o /app/build/main ./src/main.go