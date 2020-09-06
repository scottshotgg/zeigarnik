#!/bin/bash

cd docs/proto

# Look at: https://github.com/danielvladco/go-proto-gql

DOCS_DIR=../../docs

protoc \
  -I . \
  -I /usr/local/include \
  -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7/third_party/googleapis/ \
  -I ~/go/pkg/mod/github.com/ysugimoto/grpc-graphql-gateway@v0.14.8/include/graphql \
  --go_out=plugins=grpc:../../pkg \
  --grpc-gateway_out=logtostderr=true:../../pkg \
  --swagger_out=logtostderr=true:../../docs/swagger \
  --graphql_out=:../../pkg \
  *.proto

# Generate Markdown
protoc \
  -I . \
  -I /usr/local/include \
  -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7/third_party/googleapis/ \
  -I ~/go/pkg/mod/github.com/ysugimoto/grpc-graphql-gateway@v0.14.8/include/graphql \
  --doc_out=../../docs \
  --doc_opt=markdown,docs.md \
  *.proto

# Generate Markdown docs
protoc \
  -I . \
  -I /usr/local/include \
  -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7/third_party/googleapis/ \
  -I ~/go/pkg/mod/github.com/ysugimoto/grpc-graphql-gateway@v0.14.8/include/graphql \
  --doc_out=../../docs \
  --doc_opt=markdown,docs.md \
  *.proto

# Generate HTML docs
protoc \
  -I . \
  -I /usr/local/include \
  -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7/third_party/googleapis/ \
  -I ~/go/pkg/mod/github.com/ysugimoto/grpc-graphql-gateway@v0.14.8/include/graphql \
  --doc_out=../../docs \
  --doc_opt=html,docs.html \
  *.proto

# Generate JSON docs
protoc \
  -I . \
  -I /usr/local/include \
  -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7/third_party/googleapis/ \
  -I ~/go/pkg/mod/github.com/ysugimoto/grpc-graphql-gateway@v0.14.8/include/graphql \
  --doc_out=../../docs \
  --doc_opt=json,docs.json \
  *.proto

# Generate DocBook docs
protoc \
  -I . \
  -I /usr/local/include \
  -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7/third_party/googleapis/ \
  -I ~/go/pkg/mod/github.com/ysugimoto/grpc-graphql-gateway@v0.14.8/include/graphql \
  --doc_out=../../docs \
  --doc_opt=docbook,docs.xml \
  *.proto

# Generate AsciiDoc docs
protoc \
  -I . \
  -I /usr/local/include \
  -I ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7/third_party/googleapis/ \
  -I ~/go/pkg/mod/github.com/ysugimoto/grpc-graphql-gateway@v0.14.8/include/graphql \
  --doc_out=../../docs \
  --doc_opt=../../assets/templates/asciidoc.tmpl,docs.txt \
  *.proto

DOCBOOK=$DOCS_DIR/docs.xml

pandoc --from  --to epub3 --output myDocbook.epub $DOCBOOK
pandoc --from docbook --to markdown --output myDocbook.md $DOCBOOK
pandoc --from docbook --to html --output myDocbook.html $DOCBOOK
pandoc --from docbook --to latex --output myDocbook.pdf $DOCBOOK