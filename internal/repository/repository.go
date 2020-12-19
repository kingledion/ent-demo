package repository

import (
	"database/sql"
	"fmt"

	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/kingledion/ent-demo/internal/ent"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	User   string
	Pass   string
	Port   string
	DBName string
}

func New(conf Config) (*ent.Client, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s", conf.User, conf.Pass, conf.Port, conf.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)
	clnt := ent.NewClient(ent.Driver(drv))

	return clnt, nil
}
