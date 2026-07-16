package domain

import (
	"context"
	"gym-cli/internal/model/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User, userProfile *entity.UserProfile) error
}

type UserUsecase interface {
	Create(
		email, password, firstName, lastName, address string,
		memberTierId int,
	) error
}

type UserHandler interface {
	Create()
	List()
}
