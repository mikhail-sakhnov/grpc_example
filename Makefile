SRC_DIR=proto
SERVER_DST_DIR="server/proto"
CLIENT_DST_DIR="client/proto"

python:
	mkdir -p ${CLIENT_DST_DIR}
	touch ${CLIENT_DST_DIR}/__init__.py
	python -m grpc_tools.protoc -I=${SRC_DIR} --python_out=${CLIENT_DST_DIR} --grpc_python_out=${CLIENT_DST_DIR} ${SRC_DIR}/service.proto

golang:
	mkdir -p ${SERVER_DST_DIR}
	protoc -I=${SRC_DIR} --go_out=plugins=grpc:${SERVER_DST_DIR} ${SRC_DIR}/service.proto

generate: golang python


deps:
	pip install -r requirements.txt
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
