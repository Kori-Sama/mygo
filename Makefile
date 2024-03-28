NAME=mygo

build:
	@go build -o ./bin/${NAME} ./cmd/mygo
	@echo "-- Build success"

run: build
	@./bin/${NAME}

test:
	@go test -v ./test/...

clean:
	@go clean
	rm ./bin/${NAME}
	@echo "-- Clean success"