function genProto {
    DOMAIN=$1
    PROTO_PATH=./${DOMAIN}/api
    GO_OUT_PATH=./${DOMAIN}/api/gen
    mkdir -p $GO_OUT_PATH
    PBTS_OUT_DIR=./nuxt-app/service/proto_gen/${DOMAIN}

    ./tools/gen_yaml/gen_yaml -root=/Users/xxh/projects/go/safe -packageName=${DOMAIN} -tsPath=$PBTS_OUT_DIR
    protoc -I=$PROTO_PATH --grpc-gateway_out $GO_OUT_PATH --grpc-gateway_opt paths=source_relative --grpc-gateway_opt grpc_api_configuration=$PROTO_PATH/${DOMAIN}.yaml ${DOMAIN}.proto
    protoc -I=$PROTO_PATH --go_out $GO_OUT_PATH --go_opt paths=source_relative --go-grpc_out $GO_OUT_PATH --go-grpc_opt=paths=source_relative --openapiv2_out ./swagger-ui --openapiv2_opt grpc_api_configuration=$PROTO_PATH/${DOMAIN}.yaml ${DOMAIN}.proto

    mkdir -p $PBTS_OUT_DIR
    pbjs -t static -w es6 $PROTO_PATH/${DOMAIN}.proto --no-create --no-encode --no-decode --no-verify --no-delimited --force-number -o $PBTS_OUT_DIR/${DOMAIN}_pb_tmp.js
    echo 'import * as $protobuf from "protobufjs";\n' > $PBTS_OUT_DIR/${DOMAIN}_pb.js
    cat $PBTS_OUT_DIR/${DOMAIN}_pb_tmp.js >> $PBTS_OUT_DIR/${DOMAIN}_pb.js
    rm $PBTS_OUT_DIR/${DOMAIN}_pb_tmp.js
    pbts -o $PBTS_OUT_DIR/${DOMAIN}_pb.d.ts $PBTS_OUT_DIR/${DOMAIN}_pb.js
}

genProto auth
