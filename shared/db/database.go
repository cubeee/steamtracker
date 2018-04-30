package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	Db *gorm.DB
)

type ConnectDetails struct {
	Host       string
	Db         string
	User       string
	Pass       string
	Additional string
}

func ConnectPostgres(details *ConnectDetails) error {
	args := fmt.Sprintf("host=%s dbname=%s user=%s password=%s", details.Host, details.Db, details.User, details.Pass)
	if details.Additional != "" {
		args = args + " " + details.Additional
	}
	return Connect("postgres", args)
}

func Connect(dialect, args string) error {
	var err error
	Db, err = gorm.Open(dialect, args)
	Db.DB().SetMaxIdleConns(100)
	if err != nil {
		return err
	}
	return nil
}
