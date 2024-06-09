dev:
	go run main.go

	
protos:
	protoc -I protos/ protos/currency.proto  --go_out=plugins=grpc:protos/currency