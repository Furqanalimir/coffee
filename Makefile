dev:
	go run main.go
proto:
	protoc -I protos/ protos/currency.proto  --go_out=plugins=grpc:protos/currency