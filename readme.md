# map-zoo-brusque-back-go

## Como executar?

### DEV

```shell
$ go run cmd/server/main.go cmd/server/wire_gen.go
```

## Testes

Utilize o pacote [gotestsum](https://github.com/gotestyourself/gotestsum)

```shell
$ gotestsum --format=short -- -coverprofile=coverage.out ./...
$ go tool cover -html=coverage.out -o coverage.html
```