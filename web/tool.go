package web

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net/http"
	"region-cn/model"
	"region-cn/protobuf"
)

func writeJson(model interface{}, w http.ResponseWriter, indent bool) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	var bytes []byte
	var err error

	if indent {
		bytes, err = json.MarshalIndent(model, "", "  ")
	} else {
		bytes, err = json.Marshal(model)
	}

	if err != nil {
		panic(err)
	}

	if _, err = fmt.Fprint(w, string(bytes)); err != nil {
		panic(err)
	}
}

func writeProtobuf(model proto.Message, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/x-protobuf;charset=utf-8")

	data, _ := proto.Marshal(model)
	if _, err := w.Write(data); err != nil {
		panic(err)
	}

}

func mapModels(xs []model.Model) *protobuf.Models {
	var coll = &protobuf.Models{}

	for _, x := range xs {
		m := &protobuf.Model{
			Id:        x.GetId(),
			Code:      x.GetCode(),
			Name:      x.GetName(),
			ShortName: x.GetShortName(),
			Lat:       x.GetLat(),
			Lng:       x.GetLng(),
		}
		coll.List = append(coll.List, m)
	}
	return coll
}
