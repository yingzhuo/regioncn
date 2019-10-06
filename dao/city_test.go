package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestFindCityByProvinceCode(t *testing.T) {

	db := NewDB()
	defer func() {
		_ = db.Close()
	}()

	for _, it := range FindCityByProvinceCode(db, "140000") {
		fmt.Println(it)
	}

}
