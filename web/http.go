package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"region-cn/cnf"
	"region-cn/dao"
)

func StartHttpServer(db *sql.DB, config *cnf.Config) {

	// for k8s readiness checking
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {})

	http.HandleFunc("/province", func(writer http.ResponseWriter, request *http.Request) {
		xs := dao.FindAllProvinces(db)
		writeJson(xs, writer, config.IndentJson)
	})

	http.HandleFunc("/city", func(writer http.ResponseWriter, request *http.Request) {
		provinceCode := request.FormValue("provinceCode")
		xs := dao.FindCityByProvinceCode(db, provinceCode)
		writeJson(xs, writer, config.IndentJson)
	})

	http.HandleFunc("/area", func(writer http.ResponseWriter, request *http.Request) {
		cityCode := request.FormValue("cityCode")
		xs := dao.FindAreaByCityCode(db, cityCode)
		writeJson(xs, writer, config.IndentJson)
	})

	http.HandleFunc("/street", func(writer http.ResponseWriter, request *http.Request) {
		areaCode := request.FormValue("areaCode")
		xs := dao.FindStreetByAreaCode(db, areaCode)
		writeJson(xs, writer, config.IndentJson)
	})

	// start
	addr := fmt.Sprintf("%v:%v", config.HttpHost, config.HttpPort)
	_ = http.ListenAndServe(addr, nil)

}
