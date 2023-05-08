#!/usr/bin/env bash

# WARNING: This script is currently only used for generating the proto files for Tendermint Liquidity
# If we ever need to run this script for other modules, we will need to modify it a bit
# The reasons for that will be noted in other comments below
set -eo pipefail

protoc_gen_gocosmos() {
  if ! grep "github.com/gogo/protobuf => github.com/regen-network/protobuf" go.mod &>/dev/null ; then
    echo -e "\tPlease run this command from somewhere inside the cosmos-sdk folder."
    return 1
  fi

  go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos@latest 2>/dev/null
}

protoc_gen_gocosmos

proto_dirs=$(find ./proto_osmosis -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # I used v1.17.0 of buf, which has moved the protoc command under "alpha"
  # Apparently, the devs do not recommend using buf this way
  # See this issue: https://github.com/bufbuild/buf/issues/1215
  buf alpha protoc \
    -I "proto_osmosis" \
    -I "third_party/proto" \
    --gocosmos_out=plugins=interfacetype+grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. \
    --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

done

# move proto files to the right places
# We move the Tendermint protos into the tendermint subfolder
cp -r github.com/osmosis-labs/* ./osmosis/
rm -rf github.com
