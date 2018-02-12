.PHONY: install
install: protos
	go install .

.PHONY: protos
protos:
	rm -rf protos/chat.pb.go && protoc --go_out="plugins=grpc:." protos/chat.proto
