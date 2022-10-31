package auth

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authpb "micro/auth/api/gen"
	"micro/auth/dao"
	"time"
)

type Service struct {
	authpb.UnimplementedAuthServiceServer
	TokenGenerator  TokenGenerator
	PasswordManager PasswordManager
	TokenExpire     time.Duration
	Logger          *zap.Logger
	Dao             *dao.Dao
}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
}

type PasswordManager interface {
	Encode(rawPassword string) string
	Verify(rawPassword, encodedPassword string) bool
}

func (s *Service) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	record, err := s.Dao.GetPassword(request.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.Unauthenticated, "")
		}
		s.Logger.Error("", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	verify := s.PasswordManager.Verify(request.Password, record.Password)
	if !verify {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	tkn, err := s.TokenGenerator.GenerateToken(record.Id, s.TokenExpire)
	if err != nil {
		s.Logger.Error("", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		Token:       tkn,
		TokenExpire: int32(s.TokenExpire.Seconds()),
	}, nil
}
