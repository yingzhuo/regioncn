package model

import "encoding/json"

type Model interface {
	GetId() int64
	GetCode() string
	GetName() string
	GetShortName() string
	GetLat() float64
	GetLng() float64
}

// 省
type Province struct {
	// ID
	Id int64 `json:"id"`

	// 身份代码
	Code string `json:"code"`

	// 全称
	Name string `json:"name"`

	// 简称
	ShortName string `json:"shortName"`

	// 纬度
	Lat float64 `json:"lat"`

	// 经度
	Lng float64 `json:"lng"`
}

func (province *Province) GetId() int64 {
	return province.Id
}

func (province *Province) GetCode() string {
	return province.Code
}

func (province *Province) GetName() string {
	return province.ShortName
}

func (province *Province) GetShortName() string {
	return province.ShortName
}

func (province *Province) GetLat() float64 {
	return province.Lat
}

func (province *Province) GetLng() float64 {
	return province.Lng
}

func (province *Province) String() string {
	data, _ := json.Marshal(province)
	return string(data)
}

/* ------------------------------------------------------------------------------------------------------------------ */

// 市
type City struct {
	// ID
	Id int64 `json:"id"`

	// 代码
	Code string `json:"code"`

	// 全称
	Name string `json:"name"`

	// 简称
	ShortName string `json:"shortName"`

	// 纬度
	Lat float64 `json:"lat"`

	// 经度
	Lng float64 `json:"lng"`
}

func (city *City) GetId() int64 {
	return city.Id
}

func (city *City) GetCode() string {
	return city.Code
}

func (city *City) GetName() string {
	return city.Name
}

func (city *City) GetShortName() string {
	return city.ShortName
}

func (city *City) GetLat() float64 {
	return city.Lat
}

func (city *City) GetLng() float64 {
	return city.Lng
}

func (city *City) String() string {
	data, _ := json.Marshal(city)
	return string(data)
}

/* ------------------------------------------------------------------------------------------------------------------ */

// 区
type Area struct {
	// ID
	Id int64 `json:"id"`

	// 代码
	Code string `json:"code"`

	// 全称
	Name string `json:"name"`

	// 简称
	ShortName string `json:"shortName"`

	// 纬度
	Lat float64 `json:"lat"`

	// 经度
	Lng float64 `json:"lng"`
}

func (area *Area) GetId() int64 {
	return area.Id
}

func (area *Area) GetCode() string {
	return area.Code
}

func (area *Area) GetName() string {
	return area.Name
}

func (area *Area) GetShortName() string {
	return area.ShortName
}

func (area *Area) GetLat() float64 {
	return area.Lat
}

func (area *Area) GetLng() float64 {
	return area.Lng
}

func (area *Area) String() string {
	data, _ := json.Marshal(area)
	return string(data)
}

/* ------------------------------------------------------------------------------------------------------------------ */

// 街道
type Street struct {
	// ID
	Id int64 `json:"id"`

	// 代码
	Code string `json:"code"`

	// 全称
	Name string `json:"name"`

	// 简称
	ShortName string `json:"shortName"`

	// 纬度
	Lat float64 `json:"lat"`

	// 经度
	Lng float64 `json:"lng"`
}

func (street *Street) GetId() int64 {
	return street.Id
}

func (street *Street) GetCode() string {
	return street.Code
}

func (street *Street) GetName() string {
	return street.Name
}

func (street *Street) GetShortName() string {
	return street.ShortName
}

func (street *Street) GetLat() float64 {
	return street.Lat
}

func (street *Street) GetLng() float64 {
	return street.Lng
}

func (street *Street) String() string {
	data, _ := json.Marshal(street)
	return string(data)
}
