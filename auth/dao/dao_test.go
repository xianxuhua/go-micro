package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"micro/auth/pwd"
	"testing"
)

func createDao() Dao {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/micro?charset=utf8")
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	dao := Dao{
		Db:     db,
		Logger: logger,
	}
	return dao
}

func TestDao_GetAccountID(t *testing.T) {
	dao := createDao()
	record, err := dao.GetPassword("123")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(record.Password)
}

func TestDao_CreateUser(t *testing.T) {
	dao := createDao()
	manager := pwd.MD5PasswordManager{
		SaltLen:    16,
		Iterations: 100,
		KeyLen:     32,
	}
	err := dao.CreateUser("123", manager.Encode("123"))
	if err != nil {
		t.Fatal(err)
	}
}
