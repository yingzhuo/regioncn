package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestFindAllProvinces(t *testing.T) {

	db := NewDB()

	defer func() {
		_ = db.Close()
	}()

	for _, it := range FindAllProvinces(db) {
		fmt.Println(it)
	}
}

func NewDB() *sql.DB {
	dsn := "root:root@tcp(localhost:3306)/region_cn?charset=utf8mb4"
	db, _ := sql.Open("mysql", dsn)
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
