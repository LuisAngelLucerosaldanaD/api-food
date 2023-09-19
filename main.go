package main

import (
	"api-food/api"
	"api-food/internal/env"
)

// @title API FOOD
// @version 1.4
// @description Api para la gestión de la comida y recetas
// @termsOfService https://www.bjungle.net/terms/
// @contact.name API Support
// @contact.email luis.lucero@bjungle.net
// @license.name Software Owner
// @license.url https://www.bjungle.net/terms/licenses
// @host 172.147.77.149:5050
// @tag.name Category
// @tag.description Métodos referentes a las categorías de las comidas
// @tag.name Ingredient
// @tag.description Métodos referentes a los ingredientes de las comidas
// @tag.name Product
// @tag.description Métodos referentes a los productos
// @tag.name Rol
// @tag.description Métodos referentes a los roles en las recetas
// @BasePath /
func main() {
	c := env.NewConfiguration()
	api.Start(c.App.Port, c.App.ServiceName, c.App.LoggerHttp, c.App.AllowedDomains)
}
