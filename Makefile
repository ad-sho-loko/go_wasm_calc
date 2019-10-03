
build:
	GOOS=js GOARCH=wasm go build -o main.wasm

run:
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'

all: build

