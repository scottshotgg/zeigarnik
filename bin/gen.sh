#!/bin/bash

mkdir pkg/gql

cd docs/proto

# Look at: https://github.com/danielvladco/go-proto-gql

protoc -I /usr/local/include \
  -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7/third_party/googleapis/ \
  -I . \
  -I ~/go/pkg/mod/github.com/ysugimoto/grpc-graphql-gateway@v0.14.8/include/graphql \
  --go_out=plugins=grpc:../../pkg/buffs \
  --grpc-gateway_out=logtostderr=true:../../pkg/buffs \
  --swagger_out=logtostderr=true:../../docs/swagger \
  --graphql_out=:../../pkg/gql \
  *.proto

cp ../../pkg/gql/*/* ../../pkg/buffs/
rm -rf ../../pkg/gql/