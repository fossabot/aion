fmt:
	go fmt ./...
	go vet ./...
	~/go/bin/golangci-lint run 
	

test: fmt
	go test ./...

build: fmt
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o cache ./main.go 

run: build
	./cache