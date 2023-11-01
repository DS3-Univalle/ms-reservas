# Usamos la imagen oficial de Go como base
FROM golang:1.21.3

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos el código fuente de tu aplicación a la imagen
COPY . .

# Compilamos la aplicación Go
RUN go build -o main

# Exponemos el puerto en el que la aplicación escuchará
EXPOSE 4000

# Especificamos el comando para ejecutar la aplicación
CMD ["./main"]