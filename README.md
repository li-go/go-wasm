# go-wasm

#### How to run?

```
$ GOOS=js GOARCH=wasm go build -o test.wasm
$ go get -u github.com/shurcooL/goexec
$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
$ open http://localhost:8080/wasm_exec.html
```

#### How to get wasm_exec.html & wasm_exec.js

See https://github.com/golang/go/tree/master/misc/wasm
