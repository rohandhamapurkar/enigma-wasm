commands to build

```
GOOS=js GOARCH=wasm tinygo build -o main.wasm -target=wasm cmd/main.go
```

```
GOOS=js GOARCH=wasm go build -o main.wasm cmd/main.go
```