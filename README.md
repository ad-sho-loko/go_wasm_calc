# go_wasm_calc

The calculator web app written in WebAssembly. (built by Golang)

## Demo

![demo33](https://user-images.githubusercontent.com/21151388/52529075-7cf9c500-2d2e-11e9-81dc-a9ba5c08502a.gif)

## Requirement
```
$ go get -u github.com/shurcooL/goexec
```

## Build

### Windows 
```
$ set GOOS=js 
$ set GOARCH=wasm
$ go build -o main.wasm
```

### Linux
```
$ GOOS=js GOARCH=wasm go build -o main.wasm
```

## Run
```
$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```

## License

MIT

## Author

[@adsholoko](https://twitter.com/adsholoko)
