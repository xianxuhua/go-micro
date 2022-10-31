package dao

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Dao struct {
	Logger *zap.Logger
	Db     *sqlx.DB
}

type AccountIDRecord struct {
	Id       string `db:"id"`
	Password string `db:"password"`
}

func (d *Dao) GetPassword(username string) (AccountIDRecord, error) {
	record := AccountIDRecord{}
	err := d.Db.Get(&record, "select id, password from t_user where username = ?", username)
	if err != nil {
		return AccountIDRecord{}, err
	}

	return record, nil
}

func (d *Dao) CreateUser(username, password string) error {
	_, err := d.Db.Exec("insert into t_user(username, password) values (?, ?)", username, password)
	if err != nil {
		return err
	}

	return nil
}
