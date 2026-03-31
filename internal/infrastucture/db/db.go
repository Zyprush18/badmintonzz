package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Zyprush18/badmintonzz/internal/infrastucture/config"
	_ "github.com/go-sql-driver/mysql"
)


func Connect_DB() error {
	type_db,user,pass,name_db := config.GetEnv()
	db, err := sql.Open(type_db, fmt.Sprintf("%s:%s@/%s", user, pass, name_db))
	if err != nil {
		return err
	}

	defer db.Close()

	if err:= db.Ping();err != nil {
		return err
	}

	// l, err := db.Exec("IF NOT EXISTS CREATE DATABASE badmintonzz", nil)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(l)

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return nil
}
