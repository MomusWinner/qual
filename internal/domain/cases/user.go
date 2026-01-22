package cases

import (
	"app/internal/domain"
	"app/internal/domain/models"
	"context"
)

type UserUseCase struct {
	ctx domain.Context
}

func NewUserUseCase(ctx domain.Context) *UserUseCase {
	return &UserUseCase{
		ctx: ctx,
	}
}

func (uc *UserUseCase) GetById(id int) (*models.User, error) {
	user, err := uc.ctx.Connection().UserRepository().GetById(context.Background(), id)
	if err != nil {

		uc.ctx.Logger().Error(err.Error())
		return nil, ErrInternal
	}

	return user, nil
}

func (uc *UserUseCase) GetAll() (users []models.User, err error) {
	players, err := uc.ctx.Connection().UserRepository().GetAll(context.Background())
	if err != nil {
		uc.ctx.Logger().Error(err.Error())
		err = ErrInternal
		return
	}

	return players, nil
}

func (uc *UserUseCase) Create(user models.User) (*models.User, error) {
	newUser, err := uc.ctx.Connection().UserRepository().Add(context.Background(), user)
	if err != nil {
		uc.ctx.Logger().Error(err.Error())
		return nil, ErrInternal
	}

	return newUser, nil
}

func (uc *UserUseCase) Update(user models.User) (*models.User, error) {
	updatedUser, err := uc.ctx.Connection().UserRepository().Update(context.Background(), user)
	if err != nil {
		uc.ctx.Logger().Error(err.Error())
		return nil, ErrInternal
	}

	return updatedUser, nil
}

func (uc *UserUseCase) Delete(id int) error {
	err := uc.ctx.Connection().UserRepository().Delete(context.Background(), id)
	if err != nil {

		uc.ctx.Logger().Error(err.Error())
		return ErrInternal
	}

	return nil
}
