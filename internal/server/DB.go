package server

import (
	"backend_next_echo/pkg/config"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"log"
	"time"
)

var DB *bun.DB

func ConnectDb() error {
	dbConfig := config.GetDatabaseConfig()
	mysqlConfig := &mysql.Config{
		User:                 dbConfig.Username,
		Passwd:               dbConfig.Password,
		Addr:                 fmt.Sprintf("%s:%s", dbConfig.Host, dbConfig.Port),
		DBName:               dbConfig.Dbname,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
		Params: map[string]string{
			"charset":   "utf8",
			"collation": "utf8_general_ci",
		},
	}

	hsqldb, err := sql.Open(dbConfig.Connection, mysqlConfig.FormatDSN())
	if err != nil {
		return err
	}
	//defer hsqldb.Close()

	hsqldb.SetMaxOpenConns(25)
	hsqldb.SetMaxIdleConns(25)
	hsqldb.SetConnMaxLifetime(5 * time.Minute)

	DB = bun.NewDB(hsqldb, mysqldialect.New())
	//defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := DB.PingContext(ctx); err != nil {
		return err
	}
	log.Printf("connect to %s database succssfully", dbConfig.Connection)
	return nil
}

func CloseDb() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
