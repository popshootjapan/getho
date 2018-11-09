run:
	go run main.go

test:
	go test -v ./...

gobuild:
	go build -o "${GOPATH}/bin/getho" main.go
