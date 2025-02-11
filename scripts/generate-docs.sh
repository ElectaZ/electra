#!/usr/bin/env bash
# Swagger allows you to describe the structure of your APIs so that machines can read them. The ability of APIs to describe their own structure is the root of all awesomeness in Swagger.

set -eo pipefail

mkdir -p ./tmp-swagger-gen

# move the vendor folder to a temp dir so that go list works properly
temp_dir="tmpf2948f67cce"
if [ -d vendor ]; then
  mv ./vendor ./$temp_dir
fi

# Get the path of the cosmos-sdk repo from go/pkg/mod
cosmos_sdk_dir=$(go list -f '{{ .Dir }}' -m github.com/cosmos/cosmos-sdk)

# move the vendor folder back to ./vendor
if [ -d $temp_dir ]; then
  mv ./$temp_dir ./vendor
fi

proto_dirs=$(find ./proto $cosmos_sdk_dir/proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    protoc  \
    -I "proto" \
    -I "$cosmos_sdk_dir/third_party/proto" \
    -I "$cosmos_sdk_dir/proto" \
    "$query_file" \
    --swagger_out ./tmp-swagger-gen \
    --swagger_opt logtostderr=true \
    --swagger_opt fqn_for_swagger_name=true \
    --swagger_opt simple_operation_ids=true
  fi
done

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./docs/config.json -o ./docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
#rm -rf ./tmp-swagger-gen