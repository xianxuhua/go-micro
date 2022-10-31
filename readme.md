### gin + grpc + grpc-gateway + protobuf + nuxt3 + typescript
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

![img.png](https://github.com/xianxuhua/go-micro/blob/master/resources/Xnip2022-10-31_15-53-21.jpg)

![img.png](https://github.com/xianxuhua/go-micro/blob/master/resources/Xnip2022-10-31_15-53-31.jpg)

根据定义的微服务自行添加到gen.sh
```
genProto auth
genProto ...
```
6.生成rsa密钥
https://cryptotools.net/rsagen
替换public.key和private.key

7.nuxt安装
```
cd nuxt-app
yarn install
yarn dev
```
8.nuxt+nginx部署

```
nuxt generate
```
修改/etc/nginx/sites-enabled/default
```
root   xxx/nuxt-app/.output/public/;
```
修改 /etc/nginx/mime.types
```
application/javascript                           js;
```
为
```
application/javascript                           js mjs;
```
nginx -s reload

9.TODO: docker部署
