package category

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

// CreateCategory godoc
// @Summary Método que permite crear una categoría de comida
// @Description Método que permite crear una categoría de comida
// @tags Category
// @Accept json
// @Produce json
// @Param RequestCategory body RequestCategory true "Datos de la categoría"
// @Success 200 {object} ResponseCategory
// @Router /api/v1/user/category [post]
func (h handlerWork) CreateCategory(ctx *fiber.Ctx) error {
	res := ResponseCategory{Error: true}
	req := RequestCategory{}

	err := ctx.BodyParser(&req)
	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "Cuerpo de la petición invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resCategory, code, err := srvCfg.SrvCategory.CreateCategory(req.Name, req.RoleId)
	if err != nil {
		logger.Error.Println("No se pudo crear la categoría, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo crear la categoría, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Data = resCategory
	res.Code, res.Type, res.Msg = 29, 1, "Caregoría creada corractamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// UpdateCategory godoc
// @Summary Método que permite actualizar una categoría de comida
// @Description Método que permite actualizar una categoría de comida
// @tags Category
// @Accept json
// @Produce json
// @Param RequestUpdateCategory body RequestUpdateCategory true "Datos de la categoría"
// @Success 200 {object} ResponseCategory
// @Router /api/v1/user/category [put]
func (h handlerWork) UpdateCategory(ctx *fiber.Ctx) error {
	res := ResponseCategory{Error: true}
	req := RequestUpdateCategory{}

	err := ctx.BodyParser(&req)
	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "Cuerpo de la petición invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resCategory, code, err := srvCfg.SrvCategory.UpdateCategory(req.Id, req.Name, req.RoleId)
	if err != nil {
		logger.Error.Println("No se pudo actualizar la categoría, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo actualizar la categoría, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Data = resCategory
	res.Code, res.Type, res.Msg = 29, 1, "Caregoría actualizada corractamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// DeleteCategory godoc
// @Summary Método que permite borrar una categoría de comida
// @Description Método que permite borrar una categoría de comida
// @tags Category
// @Accept json
// @Produce json
// @Param id path string true "Id de la categoría"
// @Success 200 {object} ResponseAnny
// @Router /api/v1/user/category/{id} [delete]
func (h handlerWork) DeleteCategory(ctx *fiber.Ctx) error {
	res := ResponseAnny{Error: true}
	categoryId, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Println("El id de la categoría es invalido, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "El id de la categoría es invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	code, err := srvCfg.SrvCategory.DeleteCategory(categoryId)
	if err != nil {
		logger.Error.Println("No se pudo borrar la categoría, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo borrar la categoría, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Caregoría borrada corractamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// GetCategory godoc
// @Summary Método que permite obtener una categoría de comida por su id
// @Description Método que permite obtener una categoría de comida por su id
// @tags Category
// @Accept json
// @Produce json
// @Param id path string true "Id de la categoría"
// @Success 200 {object} ResponseCategory
// @Router /api/v1/user/category/{id} [get]
func (h handlerWork) GetCategory(ctx *fiber.Ctx) error {
	res := ResponseCategory{Error: true}
	categoryId, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Println("El id de la categoría es invalido, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "El id de la categoría es invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resCategory, code, err := srvCfg.SrvCategory.GetCategoryByID(categoryId)
	if err != nil {
		logger.Error.Println("No se pudo obtener la categoría, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo obtener la categoría, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	if resCategory == nil {
		res.Code, res.Type, res.Msg = 44, 3, "No existe una categoría con el id proporcionado"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resCategory
	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Procesado corractamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// GetAllCategories godoc
// @Summary Método que permite obtener todas las categorías de comida
// @Description Método que permite obtener todas las categorías de comida
// @tags Category
// @Accept json
// @Produce json
// @Success 200 {object} ResponseAllCategory
// @Router /api/v1/user/category/all [get]
func (h handlerWork) GetAllCategories(ctx *fiber.Ctx) error {
	res := ResponseAllCategory{Error: true}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resAllCategory, err := srvCfg.SrvCategory.GetAllCategory()
	if err != nil {
		logger.Error.Println("No se pudo obtener todas las categorías, err: " + err.Error())
		res.Code, res.Type, res.Msg = 22, 3, "No se pudo obtener todas las categorías, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resAllCategory
	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Procesado corractamente"
	return ctx.Status(http.StatusOK).JSON(res)
}
