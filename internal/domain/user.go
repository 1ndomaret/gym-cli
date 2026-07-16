package domain

import (
	"context"
	"gym-cli/internal/model/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User, userProfile *entity.UserProfile) error
	MemberList(ctx context.Context) ([]entity.UserDetail, error)
}

type UserUsecase interface {
	Create(
		email, password, firstName, lastName, address string,
		memberTierId int,
	) error
	MemberList() ([]entity.UserDetail, error)
}

type UserHandler interface {
	Create()
	List()
}
