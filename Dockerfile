# Establecer la imagen base de Go
FROM golang:1.20-alpine3.17

ARG APP_DIR=/api-vehicular-inspections

# Crear el directorio de trabajo
WORKDIR ${APP_DIR}

# Copiar el binario compilado
COPY main .

# Exponer el puerto 8080
EXPOSE 8080

# Iniciar la aplicación Go
CMD ["./main"]