package dao

import (
	"database/sql"
	"region-cn/model"
)

func FindCityByProvinceCode(db *sql.DB, provinceCode string) []*model.City {
	if provinceCode == "" {
		return []*model.City{}
	}

	ql := `SELECT city_id, city_code, city_name, short_name, lat, lng FROM t_region_city WHERE province_code = ?`

	rows, err := db.Query(ql, provinceCode)
	if err != nil {
		panic(err)
	}

	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	var ret []*model.City

	for rows.Next() {
		c := &model.City{}
		if err = rows.Scan(&c.Id, &c.Code, &c.Name, &c.ShortName, &c.Lat, &c.Lng); err != nil {
			panic(err)
		} else {
			ret = append(ret, c)
		}
	}

	return ret
}
