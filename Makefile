.PHONY: install
install: protos/chat.pb.go
	go install .

protos/chat.pb.go:
	protoc --go_out="plugins=grpc:." protos/chat.proto
