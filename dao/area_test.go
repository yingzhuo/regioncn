package dao

import (
	"fmt"
	"testing"
)

func TestFindAreaByCityCode(t *testing.T) {

	db := NewDB()
	defer func() {
		_ = db.Close()
	}()

	for _, it := range FindAreaByCityCode(db, "120100") {
		fmt.Println(it)
	}
}
