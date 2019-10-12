package dao

import (
	"database/sql"
	"region-cn/model"
)

func FindStreetByAreaCode(db *sql.DB, areaCode string) []model.Model {

	if areaCode == "" {
		return []model.Model{}
	}

	ql := `SELECT street_id, street_code, street_name, short_name, lat, lng FROM t_region_street WHERE area_code = ?`

	rows, err := db.Query(ql, areaCode)
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
		s := &model.Street{}
		if err := rows.Scan(&s.Id, &s.Code, &s.Name, &s.ShortName, &s.Lat, &s.Lng); err != nil {
			panic(err)
		} else {
			ret = append(ret, s)
		}
	}

	return ret
}
