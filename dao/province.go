package dao

import (
	"database/sql"
	"region-cn/model"
)

func FindAllProvinces(db *sql.DB) []model.Model {

	ql := `SELECT province_id, province_code, province_name, short_name, lat, lng FROM t_region_province`

	rows, err := db.Query(ql)
	if err != nil {
		panic(err)
	}

	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	var ret []model.Model

	for rows.Next() {
		p := &model.Province{}
		if err = rows.Scan(&p.Id, &p.Code, &p.Name, &p.ShortName, &p.Lat, &p.Lng); err != nil {
			panic(err)
		} else {
			ret = append(ret, p)
		}
	}

	return ret
}
