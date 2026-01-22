package user

import (
	"app/internal/api"
	"app/internal/domain"
	"app/internal/domain/cases"
	"app/internal/domain/models"
	"app/internal/dto"
	"app/internal/utils"
	"app/internal/validation"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	ctx         domain.Context
	userUseCase *cases.UserUseCase
}

func NewUserHandler(ctx domain.Context, userUseCase *cases.UserUseCase) *UserHandler {
	return &UserHandler{ctx, userUseCase}
}

// Create
// @Description create new user
// @Tags       User
// @Accept     json
// @Produce    json
// @Param      CreateUser body dto.CreateUser true "CreateUser"
// @Success    201 {object} dto.UserResponse
// @Failure    400 {object} api.ErrorResponse
// @Failure    409 {object} api.ErrorResponse
// @Failure    500 {object} api.ErrorResponse
// @Failure    422 {object} api.ErrorResponse
// @Router     /users [post]
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var payload *dto.CreateUser

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.NewErrorResponse([]api.Error{
			{Code: api.UnprocessableEntity, Message: err.Error()},
		}))
	}

	userErrors := validation.ValidateStruct(payload)
	if userErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse(userErrors))
	}

	hashedPassword, err := utils.HashPassword([]byte(payload.Password))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	birthday, err := time.Parse("2006-01-02", payload.Birthday)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	hasUser, _ := h.ctx.Connection().UserRepository().GetByEmail(c.Context(), payload.Email)

	if hasUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "user with the same email address already exists"})
	}

	newUser, err := h.ctx.Connection().UserRepository().Add(c.Context(), models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		Birthday: &birthday,
	})

	if err != nil {
		return api.InternalServerError(c, err, "")
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ModelToUserResponse(newUser))
}

// GetById
// @Description  Get user by id
// @Tags         User
// @Produce      json
// @Param        id path string false "user id"
// @Success      200 {object} dto.UserResponse
// @Failure      500
// @Router       /users/{id} [get]
func (h *UserHandler) GetById(c *fiber.Ctx) error {
	id, err := api.GetIdParam(c)
	if err != nil {
		return err
	}

	user, err := h.ctx.Connection().UserRepository().GetById(c.Context(), id)
	if err != nil {
		return api.InternalServerError(c, err, "")
	}

	return c.Status(fiber.StatusOK).JSON(dto.ModelToUserResponse(user))
}

// GetAll
// @Description  Get all users
// @Tags         User
// @Produce      json
// @Success      200 {object} []dto.UserResponse
// @Failure      500
// @Router       /users/ [get]
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.ctx.Connection().UserRepository().GetAll(c.Context())
	if err != nil {
		return api.InternalServerError(c, err, "")
	}

	return c.Status(fiber.StatusOK).JSON(dto.ModelsToUserResponse(users))
}

// Update
// @Description Update user by id
// @Tags       User
// @Accept     json
// @Produce    json
// @Param      id path string false "user id"
// @Param      CreateUser body dto.CreateUser true "CreateUser"
// @Success    201 {object} dto.UserResponse
// @Failure    400 {object} api.ErrorResponse
// @Failure    409 {object} api.ErrorResponse
// @Failure    500 {object} api.ErrorResponse
// @Failure    422 {object} api.ErrorResponse
// @Router     /users/{id} [put]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := api.GetIdParam(c)
	if err != nil {
		return err
	}

	var payload *dto.CreateUser

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.NewErrorResponse([]api.Error{
			{Code: api.UnprocessableEntity, Message: err.Error()},
		}))
	}

	userErrors := validation.ValidateStruct(payload)
	if userErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse(userErrors))
	}

	hashedPassword, err := utils.HashPassword([]byte(payload.Password))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	birthday, err := time.Parse("2006-01-02", payload.Birthday)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	userByEmail, _ := h.ctx.Connection().UserRepository().GetByEmail(c.Context(), payload.Email)

	if userByEmail != nil && int(userByEmail.ID) != id {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "user with the same email address already exists"})
	}

	newUser, err := h.ctx.Connection().UserRepository().Update(c.Context(), models.User{
		ID:       int32(id),
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		Birthday: &birthday,
	})

	if err != nil {
		return api.InternalServerError(c, err, "")
	}

	if newUser == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "user with your id does not exist"})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ModelToUserResponse(newUser))
}

// Delete
// @Description  Delete user by id
// @Tags         User
// @Produce      json
// @Param        id path string false "user id"
// @Success      204
// @Failure      500
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := api.GetIdParam(c)
	if err != nil {
		return err
	}

	err = h.ctx.Connection().UserRepository().Delete(c.Context(), id)
	if err != nil {
		return api.InternalServerError(c, err, "")
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"status": "success"})
}
