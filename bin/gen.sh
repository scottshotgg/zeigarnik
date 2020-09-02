#!/bin/bash

cd docs/proto

protoc -I /usr/local/include \
  -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7/third_party/googleapis/ \
  -I . \
  --go_out=plugins=grpc:../../pkg/buffs \
  --grpc-gateway_out=logtostderr=true:../../pkg/buffs \
  --swagger_out=logtostderr=true:../../docs/swagger \
  *.proto