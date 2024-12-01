# api-vehicular-inspections

## Información

### Versión

0.0.1

### Descripción

g1 app

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
ó
```bash
make run-test
```

## Endpoints de estado de salud

**Liveness:** Implementado por defecto en la URL:

> localhost:8080/liveness

**Readiness:** Implementar en cmd/server/server.go

> localhost:8080/readiness
