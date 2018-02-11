.PHONY: install
install: protos
	go install .

.PHONY: protos
protos:
	rm -rf protos/echo.pb.go && protoc --go_out="plugins=grpc:." protos/echo.proto
