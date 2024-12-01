# api-vehicular-inspections

## Información

### Versión

0.0.1

### Descripción

api-vehicular-inspections

### Autor

hidalgo

## Compilación

```bash
go build -o main ./cmd
```

## Ejecución

```bash
./main
```

## Ejecución en entorno local

```bash
go run cmd/main.go
```

## Endpoints de estado de salud

**Liveness:** Implementado por defecto en la URL:

> localhost:8080/liveness

**Readiness:** Implementar en cmd/server/server.go

> localhost:8080/readiness
