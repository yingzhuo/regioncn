package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"region-cn/cnf"
	"region-cn/dao"
	"strings"
)

func StartHttpServer(db *sql.DB, config *cnf.Config) {

	// for k8s readiness checking
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {})

	http.HandleFunc("/province", func(writer http.ResponseWriter, request *http.Request) {
		xs := dao.FindAllProvinces(db)

		if strings.EqualFold(cnf.Json, config.ResponseType) {
			writeJson(xs, writer, config.IndentJson)
		} else {
			writeProtobuf(mapModels(xs), writer)
		}
	})

	http.HandleFunc("/city", func(writer http.ResponseWriter, request *http.Request) {
		provinceCode := request.FormValue("provinceCode")
		xs := dao.FindCityByProvinceCode(db, provinceCode)

		if strings.EqualFold(cnf.Json, config.ResponseType) {
			writeJson(xs, writer, config.IndentJson)
		} else {
			writeProtobuf(mapModels(xs), writer)
		}
	})

	http.HandleFunc("/area", func(writer http.ResponseWriter, request *http.Request) {
		cityCode := request.FormValue("cityCode")
		xs := dao.FindAreaByCityCode(db, cityCode)

		if strings.EqualFold(cnf.Json, config.ResponseType) {
			writeJson(xs, writer, config.IndentJson)
		} else {
			writeProtobuf(mapModels(xs), writer)
		}
	})

	http.HandleFunc("/street", func(writer http.ResponseWriter, request *http.Request) {
		areaCode := request.FormValue("areaCode")
		xs := dao.FindStreetByAreaCode(db, areaCode)

		if strings.EqualFold(cnf.Json, config.ResponseType) {
			writeJson(xs, writer, config.IndentJson)
		} else {
			writeProtobuf(mapModels(xs), writer)
		}
	})

	// start
	addr := fmt.Sprintf("0.0.0.0:%v", config.HttpPort)
	_ = http.ListenAndServe(addr, nil)

}
