#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)

GEN_OUT_DIR=${CUR}/gen

GEN_TOOL=${CUR}/protoc_gen_plugin.bash
PROTO_PATH=${CUR}/idl
PROTO_INCLUDE_PATH_1="${CUR}/idl"

check_flag_value_set() {
  if [ -z "${1}" ]; then
    exit 1
  fi
}

# ==============================================================
# LANGUAGE - Go
LANG_OUT=${GEN_OUT_DIR}/go
check_flag_value_set "${GO_MODULE_NAME}"

# go
${GEN_TOOL} \
  --proto_path=${PROTO_PATH} \
  --proto_include_path=${PROTO_INCLUDE_PATH_1} \
  --plugin_name=go \
  --plugin_out=${LANG_OUT} \
  --plugin_opt=paths=source_relative
# grpc
${GEN_TOOL} \
  --proto_path=${PROTO_PATH} \
  --proto_include_path=${PROTO_INCLUDE_PATH_1} \
  --plugin_name=go-grpc \
  --plugin_out=${LANG_OUT} \
  --plugin_opt=paths=source_relative
# go-micro
${GEN_TOOL} \
  --proto_path=${PROTO_PATH} \
  --proto_include_path=${PROTO_INCLUDE_PATH_1} \
  --plugin_name=micro \
  --plugin_out=${LANG_OUT} \
  --plugin_opt=paths=source_relative

# initialize go.mod
pushd ${LANG_OUT}
#go mod init ${GO_MODULE_NAME}
# FIXME: etcd 与 grpc 1.30+ 版本不兼容，此处降级到 1.29
go mod edit --replace=google.golang.org/grpc=google.golang.org/grpc@v1.29.1
find . -type f -name '*_grpc.pb.go' | xargs -I {} sed -i '' -e 's/SupportPackageIsVersion7/SupportPackageIsVersion6/g' {}

go build -o /dev/null ./...
popd

# ==============================================================
# LANGUAGE - TypeScript (node)
LANG_OUT=${GEN_OUT_DIR}/ts-node
${GEN_TOOL} \
  --proto_path=${PROTO_PATH} \
  --proto_include_path=${PROTO_INCLUDE_PATH_1} \
  --plugin_name=ts \
  --plugin_out=${LANG_OUT} \
  --plugin_opt=paths=source_relative,service=grpc-node

# ==============================================================
# LANGUAGE - JavaScript (commonjs + grpc-web)
LANG_OUT=${GEN_OUT_DIR}/js-grpc-web
#   commonjs
${GEN_TOOL} \
  --proto_path=${PROTO_PATH} \
  --proto_include_path=${PROTO_INCLUDE_PATH_1} \
  --plugin_name=js \
  --plugin_out=${LANG_OUT} \
  --plugin_out_opt=import_style=commonjs
#   grpcwebtext
${GEN_TOOL} \
  --proto_path=${PROTO_PATH} \
  --proto_include_path=${PROTO_INCLUDE_PATH_1} \
  --plugin_name=grpc-web \
  --plugin_out=${LANG_OUT} \
  --plugin_out_opt=import_style=commonjs+dts,mode=grpcwebtext

# ==============================================================
# LANGUAGE - Java
LANG_OUT=${GEN_OUT_DIR}/java
${GEN_TOOL} \
  --proto_path=${PROTO_PATH} \
  --proto_include_path=${PROTO_INCLUDE_PATH_1} \
  --plugin_name=java \
  --plugin_out=${LANG_OUT}

# ==============================================================
# LANGUAGE - Swift for iOS
# LANG_OUT=${GEN_OUT_DIR}/swift-ios
# # swift pb
# ${GEN_TOOL} \
#   --proto_path=${PROTO_PATH} \
#   --proto_include_path=${PROTO_INCLUDE_PATH_1} \
#   --plugin_name=swift \
#   --plugin_out=${LANG_OUT}
# # swfit grpc
# ${GEN_TOOL} \
#   --proto_path=${PROTO_PATH} \
#   --proto_include_path=${PROTO_INCLUDE_PATH_1} \
#   --plugin_name=grpc-swift \
#   --plugin_out=${LANG_OUT}
