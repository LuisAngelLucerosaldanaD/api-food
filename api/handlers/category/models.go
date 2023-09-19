package category

import "api-food/pkg/cfg/category"

type ResponseCategory struct {
	Error bool               `json:"error"`
	Data  *category.Category `json:"data"`
	Code  int                `json:"code"`
	Type  int                `json:"type"`
	Msg   string             `json:"msg"`
}

type ResponseAnny struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
	Type  int         `json:"type"`
	Msg   string      `json:"msg"`
}

type ResponseAllCategory struct {
	Error bool                 `json:"error"`
	Data  []*category.Category `json:"data"`
	Code  int                  `json:"code"`
	Type  int                  `json:"type"`
	Msg   string               `json:"msg"`
}

type RequestCategory struct {
	Name   string `json:"name"`
	RoleId int64  `json:"role_id"`
}

type RequestUpdateCategory struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	RoleId int64  `json:"role_id"`
}
