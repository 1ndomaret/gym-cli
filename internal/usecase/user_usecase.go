package usecase

import (
	"context"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
	"time"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

var setTimeout = 3 * time.Millisecond

func (u *userUsecase) Create(
	email, password, firstName, lastName, address string,
	memberTierId int,
) error {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, setTimeout)
	defer cancel()

	user := &entity.User{
		Email:    email,
		Password: password,
	}

	userProfile := &entity.UserProfile{
		MemberTierId: memberTierId,
		FirstName:    firstName,
		LastName:     lastName,
		Address:      address,
	}

	return u.repo.Create(ctx, user, userProfile)
}

func (u *userUsecase) MemberList() ([]entity.UserDetail, error) {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, setTimeout)
	defer cancel()
	return u.repo.MemberList(ctx)
}
