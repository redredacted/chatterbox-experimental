run:
	go run main.go

compile:

	go build -o bin/main main.go

clean:

	rm -R bin

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./protobuf/*.proto

# `.PHONY:` is used to differentiate files from commands
.PHONY: run
.PHONY: compile
.PHONY: clean
.PHONY: protoc