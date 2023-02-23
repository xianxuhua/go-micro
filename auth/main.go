package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	authpb "micro/auth/api/gen"
	"micro/auth/auth"
	"micro/auth/dao"
	"micro/auth/pwd"
	"micro/auth/token"
	"micro/shared/addr"
	"net"
	"os"
	"time"
)

var port = flag.Int("port", 8888, "")
var debug = flag.Bool("debug", false, "")
var dsn = flag.String("dsn", "root:123456@tcp(127.0.0.1:3306)/safe_calc?charset=utf8", "")

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	if *debug == false {
		freePort, err := addr.FreePort()
		if err != nil {
			logger.Fatal("cannot get free port", zap.Error(err))
		}
		*port = freePort
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}

	server := grpc.NewServer()
	db, err := sqlx.Connect("mysql", *dsn)
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}

	pkFile, err := os.Open("auth/private.key")
	if err != nil {
		logger.Fatal("can not open private.key", zap.Error(err))
	}
	pkBytes, err := io.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("can not read private.key", zap.Error(err))
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("can not parse private.key", zap.Error(err))
	}
	authpb.RegisterAuthServiceServer(server, &auth.Service{
		TokenGenerator: &token.JWTTokenGen{
			Issuer:     "xxh",
			NowFunc:    time.Now,
			PrivateKey: privateKey,
		},
		PasswordManager: &pwd.MD5PasswordManager{
			SaltLen:    16,
			Iterations: 100,
			KeyLen:     32,
		},
		TokenExpire: time.Hour * 24 * 30,
		Logger:      logger,
		Dao: &dao.Dao{
			Logger: logger,
			Db:     db,
		},
	})
	err = server.Serve(lis)
	if err != nil {
		logger.Fatal("cannot listen server", zap.Error(err))
	}
}
