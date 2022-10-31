### gin + grpc + protobuf + nuxt + typescript
来源于ccmouse老师

无需接口硬编码

1.安装protobuf
```
brew install protobuf
```
2.安装grpc-gateway
```
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
3.编译tools/gen_yaml/main.go

作用：
生成前端访问URL静态定义、
生成后端yaml（后端访问URL）
```
go build -o gen_yaml main.go
```
4.运行gen.sh

作用：生成前后端接口定义、URL定义

根据定义的微服务自行添加到gen.sh
```
genProto auth
```
5.nuxt安装
```
cd nuxt-app
yarn install
yarn dev
```