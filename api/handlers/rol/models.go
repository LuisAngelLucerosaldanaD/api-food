package rol

import "api-food/pkg/cfg/rol"

type ResponseRol struct {
	Error bool     `json:"error"`
	Data  *rol.Rol `json:"data"`
	Code  int      `json:"code"`
	Type  int      `json:"type"`
	Msg   string   `json:"msg"`
}

type ResponseAnny struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
	Type  int         `json:"type"`
	Msg   string      `json:"msg"`
}

type ResponseAllRol struct {
	Error bool       `json:"error"`
	Data  []*rol.Rol `json:"data"`
	Code  int        `json:"code"`
	Type  int        `json:"type"`
	Msg   string     `json:"msg"`
}

type RequestRol struct {
	Name string `json:"name"`
}

type RequestUpdateRol struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
