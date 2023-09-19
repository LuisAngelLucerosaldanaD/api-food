package cfg

import (
	"api-food/pkg/cfg/category"
	"api-food/pkg/cfg/ingredient"
	"api-food/pkg/cfg/product"
	"api-food/pkg/cfg/rol"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvCategory   category.PortsServerCategory
	SrvIngredient ingredient.PortsServerIngredient
	SrvProduct    product.PortsServerProduct
	SrvRol        rol.PortsServerRol
}

func NewServerCfg(db *sqlx.DB, txID string) *Server {

	repoCategory := category.FactoryStorage(db, txID)
	srvCategory := category.NewCategoryService(repoCategory, txID)

	repoIngredient := ingredient.FactoryStorage(db, txID)
	srvIngredient := ingredient.NewIngredientService(repoIngredient, txID)

	repoProduct := product.FactoryStorage(db, txID)
	srvProduct := product.NewProductService(repoProduct, txID)

	repoRol := rol.FactoryStorage(db, txID)
	srvRol := rol.NewRolService(repoRol, txID)

	return &Server{
		SrvCategory:   srvCategory,
		SrvIngredient: srvIngredient,
		SrvProduct:    srvProduct,
		SrvRol:        srvRol,
	}
}
