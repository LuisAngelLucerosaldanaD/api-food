package ingredient

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

// CreateIngredient godoc
// @Summary Método que permite crear un ingrediente de comida
// @Description Método que permite crear un ingrediente de comida
// @tags Ingredient
// @Accept json
// @Produce json
// @Param RequestIngredient body RequestIngredient true "Datos del ingrediente"
// @Success 200 {object} ResponseIngredient
// @Router /api/v1/user/ingredient [post]
func (h handlerWork) CreateIngredient(ctx *fiber.Ctx) error {
	res := ResponseIngredient{Error: true}
	req := RequestIngredient{}

	err := ctx.BodyParser(&req)
	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "Cuerpo de la petición invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resIngredient, code, err := srvCfg.SrvIngredient.CreateIngredient(req.Name, req.IdProduct)
	if err != nil {
		logger.Error.Println("No se pudo crear el ingrediente, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo crear el ingrediente, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Data = resIngredient
	res.Code, res.Type, res.Msg = 29, 1, "Ingredient creada correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// UpdateIngredient godoc
// @Summary Método que permite actualizar un ingrediente de comida
// @Description Método que permite actualizar un ingrediente de comida
// @tags Ingredient
// @Accept json
// @Produce json
// @Param RequestUpdateIngredient body RequestUpdateIngredient true "Datos del ingrediente"
// @Success 200 {object} ResponseIngredient
// @Router /api/v1/user/ingredient [put]
func (h handlerWork) UpdateIngredient(ctx *fiber.Ctx) error {
	res := ResponseIngredient{Error: true}
	req := RequestUpdateIngredient{}

	err := ctx.BodyParser(&req)
	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "Cuerpo de la petición invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resIngredient, code, err := srvCfg.SrvIngredient.UpdateIngredient(req.Id, req.Name, req.IdProduct)
	if err != nil {
		logger.Error.Println("No se pudo actualizar el ingrediente, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo actualizar el ingrediente, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Data = resIngredient
	res.Code, res.Type, res.Msg = 29, 1, "Ingredient actualizada correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// DeleteIngredient godoc
// @Summary Método que permite borrar un ingrediente de comida
// @Description Método que permite borrar un ingrediente de comida
// @tags Ingredient
// @Accept json
// @Produce json
// @Param id path string true "Id del ingrediente"
// @Success 200 {object} ResponseAnny
// @Router /api/v1/user/ingredient/{id} [delete]
func (h handlerWork) DeleteIngredient(ctx *fiber.Ctx) error {
	res := ResponseAnny{Error: true}
	IngredientId, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Println("El id del ingrediente es invalido, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "El id del ingrediente es invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	code, err := srvCfg.SrvIngredient.DeleteIngredient(IngredientId)
	if err != nil {
		logger.Error.Println("No se pudo borrar el ingrediente, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo borrar el ingrediente, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Ingrediente borrada correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// GetIngredient godoc
// @Summary Método que permite obtener un ingrediente de comida por su id
// @Description Método que permite obtener un ingrediente de comida por su id
// @tags Ingredient
// @Accept json
// @Produce json
// @Param id path string true "Id del ingrediente"
// @Success 200 {object} ResponseIngredient
// @Router /api/v1/user/ingredient/{id} [get]
func (h handlerWork) GetIngredient(ctx *fiber.Ctx) error {
	res := ResponseIngredient{Error: true}
	IngredientId, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Println("El id del ingrediente es invalido, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "El id del ingrediente es invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resIngredient, code, err := srvCfg.SrvIngredient.GetIngredientByID(IngredientId)
	if err != nil {
		logger.Error.Println("No se pudo obtener la ingrediente, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo obtener la ingrediente, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	if resIngredient == nil {
		res.Code, res.Type, res.Msg = 44, 3, "No existe una ingrediente con el id proporcionado"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resIngredient
	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Procesado correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// GetAllIngredients godoc
// @Summary Método que permite obtener todos los ingredientes de comida
// @Description Método que permite obtener todos los ingredientes de comida
// @tags Ingredient
// @Accept json
// @Produce json
// @Success 200 {object} ResponseAllIngredients
// @Router /api/v1/user/ingredient/all [get]
func (h handlerWork) GetAllIngredients(ctx *fiber.Ctx) error {
	res := ResponseAllIngredients{Error: true}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resAllIngredient, err := srvCfg.SrvIngredient.GetAllIngredient()
	if err != nil {
		logger.Error.Println("No se pudo obtener todas los ingredientes, err: " + err.Error())
		res.Code, res.Type, res.Msg = 22, 3, "No se pudo obtener todas los ingredientes, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resAllIngredient
	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Procesado correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}
