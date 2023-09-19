package product

import (
	"api-food/internal/logger"
	"api-food/pkg/cfg"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type handlerWork struct {
	DB   *sqlx.DB
	TxID string
}

// CreateProduct godoc
// @Summary Método que permite crear un producto de comida
// @Description Método que permite crear un producto de comida
// @tags Product
// @Accept json
// @Produce json
// @Param RequestProduct body RequestProduct true "Datos de el producto"
// @Success 200 {object} ResponseProduct
// @Router /api/v1/user/product [post]
func (h handlerWork) CreateProduct(ctx *fiber.Ctx) error {
	res := ResponseProduct{Error: true}
	req := RequestProduct{}

	err := ctx.BodyParser(&req)
	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "Cuerpo de la petición invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resProduct, code, err := srvCfg.SrvProduct.CreateProduct(req.Name, req.UrlImg, req.Time, req.Description, req.CategoryId, req.Calorie)
	if err != nil {
		logger.Error.Println("No se pudo crear el producto, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo crear el producto, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Data = resProduct
	res.Code, res.Type, res.Msg = 29, 1, "Producto creada correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// UpdateProduct godoc
// @Summary Método que permite actualizar un producto de comida
// @Description Método que permite actualizar un producto de comida
// @tags Product
// @Accept json
// @Produce json
// @Param RequestUpdateProduct body RequestUpdateProduct true "Datos del producto"
// @Success 200 {object} ResponseProduct
// @Router /api/v1/user/product [put]
func (h handlerWork) UpdateProduct(ctx *fiber.Ctx) error {
	res := ResponseProduct{Error: true}
	req := RequestUpdateProduct{}

	err := ctx.BodyParser(&req)
	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "Cuerpo de la petición invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resProduct, code, err := srvCfg.SrvProduct.UpdateProduct(req.Id, req.Name, req.UrlImg, req.Time, req.Description, req.CategoryId, req.Calorie)
	if err != nil {
		logger.Error.Println("No se pudo actualizar el producto, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo actualizar el producto, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Data = resProduct
	res.Code, res.Type, res.Msg = 29, 1, "Producto actualizada correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// DeleteProduct godoc
// @Summary Método que permite borrar un producto de comida
// @Description Método que permite borrar un producto de comida
// @tags Product
// @Accept json
// @Produce json
// @Param id path string true "Id del producto"
// @Success 200 {object} ResponseAnny
// @Router /api/v1/user/product/{id} [delete]
func (h handlerWork) DeleteProduct(ctx *fiber.Ctx) error {
	res := ResponseAnny{Error: true}
	ProductId, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Println("El id del producto es invalido, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "El id de el producto es invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	code, err := srvCfg.SrvProduct.DeleteProduct(ProductId)
	if err != nil {
		logger.Error.Println("No se pudo borrar el producto, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo borrar el producto, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Producto borrada correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// GetProduct godoc
// @Summary Método que permite obtener un producto de comida por su id
// @Description Método que permite obtener un producto de comida por su id
// @tags Product
// @Accept json
// @Produce json
// @Param id path string true "Id del producto"
// @Success 200 {object} ResponseProduct
// @Router /api/v1/user/product/{id} [get]
func (h handlerWork) GetProduct(ctx *fiber.Ctx) error {
	res := ResponseProduct{Error: true}
	ProductId, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Println("El id de el producto es invalido, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "El id del producto es invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resProduct, code, err := srvCfg.SrvProduct.GetProductByID(ProductId)
	if err != nil {
		logger.Error.Println("No se pudo obtener el producto, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo obtener el producto, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	if resProduct == nil {
		res.Code, res.Type, res.Msg = 44, 3, "No existe un producto con el id proporcionado"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resProduct
	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Procesado correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// GetAllProducts godoc
// @Summary Método que permite obtener todos los productos de comida
// @Description Método que permite obtener todos los productos de comida
// @tags Product
// @Accept json
// @Produce json
// @Success 200 {object} ResponseAllProducts
// @Router /api/v1/user/product/all [get]
func (h handlerWork) GetAllProducts(ctx *fiber.Ctx) error {
	res := ResponseAllProducts{Error: true}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resAllProduct, err := srvCfg.SrvProduct.GetAllProduct()
	if err != nil {
		logger.Error.Println("No se pudo obtener todos los productos, err: " + err.Error())
		res.Code, res.Type, res.Msg = 22, 3, "No se pudo obtener todos los productos, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resAllProduct
	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Procesado correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}
