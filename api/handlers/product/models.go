package product

import "api-food/pkg/cfg/product"

type ResponseProduct struct {
	Error bool             `json:"error"`
	Data  *product.Product `json:"data"`
	Code  int              `json:"code"`
	Type  int              `json:"type"`
	Msg   string           `json:"msg"`
}

type ResponseAnny struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
	Type  int         `json:"type"`
	Msg   string      `json:"msg"`
}

type ResponseAllProducts struct {
	Error bool               `json:"error"`
	Data  []*product.Product `json:"data"`
	Code  int                `json:"code"`
	Type  int                `json:"type"`
	Msg   string             `json:"msg"`
}

type RequestProduct struct {
	Name        string `json:"name"`
	UrlImg      string `json:"url_img"`
	Time        string `json:"time"`
	Description string `json:"description"`
	CategoryId  int64  `json:"category_id"`
	Calorie     string `json:"calorie"`
}

type RequestUpdateProduct struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	UrlImg      string `json:"url_img"`
	Time        string `json:"time"`
	Description string `json:"description"`
	CategoryId  int64  `json:"category_id"`
	Calorie     string `json:"calorie"`
}
