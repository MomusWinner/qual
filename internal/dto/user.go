package dto

import (
	"app/internal/domain/models"
	"time"
)

type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Birthday string `json:"birthday" validate:"required"`
}

type UserResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
}

func ModelToUserResponse(m *models.User) UserResponse {
	return UserResponse{
		Id:        int(m.ID),
		Name:      m.Name,
		Email:     m.Email,
		Birthday:  *m.Birthday,
		CreatedAt: *m.CreatedAt,
	}
}

func ModelsToUserResponse(m []models.User) []UserResponse {
	users := make([]UserResponse, len(m))
	for i := range m {
		users[i] = ModelToUserResponse(&m[i])
	}

	return users
}
