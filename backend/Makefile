build:
	cd cmd && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ../build/ormuco

run:
	@./build/ormuco

test:
	@go test -v -cover ./...

clean:
	rm -rf build