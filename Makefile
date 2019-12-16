fmt:
	go fmt ./...
	go vet ./...
	~/go/bin/golangci-lint run -D errcheck
	

test: fmt
	go test -covermode=atomic ./...

build: fmt
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o aion ./main.go 

benchmark: fmt
	go test -bench=. -benchmem -benchtime=4s ./... -timeout 30m

run: build
	./aion