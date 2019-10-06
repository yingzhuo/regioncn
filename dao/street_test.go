package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestFindStreetByAreaCode(t *testing.T) {

	db := NewDB()
	defer func() {
		_ = db.Close()
	}()

	for _, it := range FindStreetByAreaCode(db, "110109") {
		fmt.Println(it)
	}

}
