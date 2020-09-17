# schutterij
Backend and middleware component

## Build locally
```
go build -o ./build ./cmd/main.go
```

## Build container image
```
docker build -t jordan396/schutterij .
docker run jordan396/schutterij
docker push jordan396/schutterij
```