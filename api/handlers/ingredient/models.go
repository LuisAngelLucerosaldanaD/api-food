package ingredient

import "api-food/pkg/cfg/ingredient"

type ResponseIngredient struct {
	Error bool                   `json:"error"`
	Data  *ingredient.Ingredient `json:"data"`
	Code  int                    `json:"code"`
	Type  int                    `json:"type"`
	Msg   string                 `json:"msg"`
}

type ResponseAnny struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
	Type  int         `json:"type"`
	Msg   string      `json:"msg"`
}

type ResponseAllIngredients struct {
	Error bool                     `json:"error"`
	Data  []*ingredient.Ingredient `json:"data"`
	Code  int                      `json:"code"`
	Type  int                      `json:"type"`
	Msg   string                   `json:"msg"`
}

type RequestIngredient struct {
	Name      string `json:"name"`
	IdProduct int64  `json:"id_product"`
}

type RequestUpdateIngredient struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	IdProduct int64  `json:"id_product"`
}
