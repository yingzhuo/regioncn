package dao

import (
	"database/sql"
	"region-cn/model"
)

func FindAreaByCityCode(db *sql.DB, cityCode string) []model.Model {

	if cityCode == "" {
		return []model.Model{}
	}

	ql := `SELECT area_id, area_code, area_name, short_name, lat, lng FROM t_region_area WHERE city_code = ?`

	rows, err := db.Query(ql, cityCode)
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
		a := &model.Area{}
		if err = rows.Scan(&a.Id, &a.Code, &a.Name, &a.ShortName, &a.Lat, &a.Lng); err != nil {
			panic(err)
		} else {
			ret = append(ret, a)
		}
	}

	return ret
}
