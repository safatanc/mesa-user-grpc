generate:
	protoc --go_out=pb/user --go_opt=paths=source_relative --go-grpc_out=pb/user --go-grpc_opt=paths=source_relative proto/user.proto

run:
	go run server.go