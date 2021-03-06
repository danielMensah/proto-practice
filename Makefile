BIN = bin/proto-go-course
PROTO_DIR = proto

UNAME := $(shell uname -s)
OS = macos
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

build: 	generate
	go build -o ${BIN} .

generate:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. ${PROTO_DIR}/*.proto

clean:
	rm ${PROTO_DIR}/*.pb.go
	rm ${BIN}
