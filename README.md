# Why go

- Fast compilation
- Cross compilation
- No separate runtime installation
- Static Binaries
- Statically typed
- Compiled Programming Language
- Simple Language (25 keywords)
- Better fit for Cloud Native Computing 
- Low Latency GC (1-10 mils)
- High Performance
- Fast Startup times
- Concurrency (native support)
- No need of any external webserver
- Systems programming
- Statically Linked 

- list of cross compilations

```go tool dist list
```

- cross compilation and build

```
GOOS=linux GOARCH=amd64 go build -o build/demo-linux-64.exe main.go
```

- release build (stripe down build)

```
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/demo-release-linux-64.exe main.go
```

## compilation proce

```
go tool compile -o demo.o main.go
```

```
go tool link -o demo demo.o
```