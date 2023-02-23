package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	authpb "micro/auth/api/gen"
	"micro/middleware"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers: true,
			},
		},
	))

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	services := []struct {
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
		addr         string
	}{
		{
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
			addr:         "localhost:8888",
		},
	}
	for _, service := range services {
		err := service.registerFunc(ctx, mux, service.addr, []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	engine := gin.Default()
	//engine.Use(middleware.Auth())
	engine.Use(middleware.Cors())
	engine.POST("/*method", gin.WrapH(mux))
	engine.StaticFS("/swagger-ui/", http.Dir("./swagger-ui"))

	err := engine.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
}
