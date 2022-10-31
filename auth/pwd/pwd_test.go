package pwd

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"micro/auth/dao"
	"testing"
)

func TestMD5PasswordManager_Encode(t *testing.T) {
	passwordManager := MD5PasswordManager{
		SaltLen:    16,
		Iterations: 100,
		KeyLen:     32,
	}
	t.Log(passwordManager.Encode("123"))
}

func TestMD5PasswordManager_Verify(t *testing.T) {
	passwordManager := MD5PasswordManager{
		SaltLen:    16,
		Iterations: 100,
		KeyLen:     32,
	}
	if !passwordManager.Verify("123", "$pbkdf2-sha512$OWh25DfLMAvtptIb$041b7756cf21360dc003eed63e60b4ed04d27d07bc3d91abe02433a75f90bd80") {
		t.Fatal("verify failed")
	}

}

func createDao() dao.Dao {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/micro?charset=utf8")
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return dao.Dao{
		Db:     db,
		Logger: logger,
	}
}
