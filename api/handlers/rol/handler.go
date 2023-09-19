package rol

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

// CreateRol godoc
// @Summary Método que permite crear un rol de comida
// @Description Método que permite crear un rol de comida
// @tags Rol
// @Accept json
// @Produce json
// @Param RequestRol body RequestRol true "Datos del rol"
// @Success 200 {object} ResponseRol
// @Router /api/v1/user/rol [post]
func (h handlerWork) CreateRol(ctx *fiber.Ctx) error {
	res := ResponseRol{Error: true}
	req := RequestRol{}

	err := ctx.BodyParser(&req)
	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "Cuerpo de la petición invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resRol, code, err := srvCfg.SrvRol.CreateRol(req.Name)
	if err != nil {
		logger.Error.Println("No se pudo crear el rol, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo crear el rol, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Data = resRol
	res.Code, res.Type, res.Msg = 29, 1, "Rol creada correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// UpdateRol godoc
// @Summary Método que permite actualizar un rol de comida
// @Description Método que permite actualizar un rol de comida
// @tags Rol
// @Accept json
// @Produce json
// @Param RequestUpdateRol body RequestUpdateRol true "Datos del rol"
// @Success 200 {object} ResponseRol
// @Router /api/v1/user/rol [put]
func (h handlerWork) UpdateRol(ctx *fiber.Ctx) error {
	res := ResponseRol{Error: true}
	req := RequestUpdateRol{}

	err := ctx.BodyParser(&req)
	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "Cuerpo de la petición invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resRol, code, err := srvCfg.SrvRol.UpdateRol(req.Id, req.Name)
	if err != nil {
		logger.Error.Println("No se pudo actualizar el rol, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo actualizar el rol, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Data = resRol
	res.Code, res.Type, res.Msg = 29, 1, "Rol actualizado correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// DeleteRol godoc
// @Summary Método que permite borrar un rol de comida
// @Description Método que permite borrar un rol de comida
// @tags Rol
// @Accept json
// @Produce json
// @Param id path string true "Id del rol"
// @Success 200 {object} ResponseAnny
// @Router /api/v1/user/rol/{id} [delete]
func (h handlerWork) DeleteRol(ctx *fiber.Ctx) error {
	res := ResponseAnny{Error: true}
	RolId, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Println("El id del rol es invalido, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "El id del rol es invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	code, err := srvCfg.SrvRol.DeleteRol(RolId)
	if err != nil {
		logger.Error.Println("No se pudo borrar el rol, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo borrar el rol, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Rol borrado correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// GetRol godoc
// @Summary Método que permite obtener un rol de comida por su id
// @Description Método que permite obtener un rol de comida por su id
// @tags Rol
// @Accept json
// @Produce json
// @Param id path string true "Id del rol"
// @Success 200 {object} ResponseRol
// @Router /api/v1/user/rol/{id} [get]
func (h handlerWork) GetRol(ctx *fiber.Ctx) error {
	res := ResponseRol{Error: true}
	RolId, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Println("El id del rol es invalido, error: " + err.Error())
		res.Code, res.Type, res.Msg = 15, 3, "El id del rol es invalido"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resRol, code, err := srvCfg.SrvRol.GetRolByID(RolId)
	if err != nil {
		logger.Error.Println("No se pudo obtener el rol, err: " + err.Error())
		res.Code, res.Type, res.Msg = code, 3, "No se pudo obtener el rol, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	if resRol == nil {
		res.Code, res.Type, res.Msg = 44, 3, "No existe un rol con el id proporcionado"
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resRol
	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Procesado correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}

// GetAllRoles godoc
// @Summary Método que permite obtener todas las categorías de comida
// @Description Método que permite obtener todas las categorías de comida
// @tags Rol
// @Accept json
// @Produce json
// @Success 200 {object} ResponseAllRol
// @Router /api/v1/user/rol/all [get]
func (h handlerWork) GetAllRoles(ctx *fiber.Ctx) error {
	res := ResponseAllRol{Error: true}

	srvCfg := cfg.NewServerCfg(h.DB, h.TxID)
	resAllRol, err := srvCfg.SrvRol.GetAllRol()
	if err != nil {
		logger.Error.Println("No se pudo obtener todas las categorías, err: " + err.Error())
		res.Code, res.Type, res.Msg = 22, 3, "No se pudo obtener todas las categorías, err: "+err.Error()
		return ctx.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = resAllRol
	res.Error = false
	res.Code, res.Type, res.Msg = 29, 1, "Procesado correctamente"
	return ctx.Status(http.StatusOK).JSON(res)
}
