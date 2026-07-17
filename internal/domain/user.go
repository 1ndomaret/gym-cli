package domain

import (
	"context"
	"gym-cli/internal/model/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User, userProfile *entity.UserProfile) error
	MemberList(ctx context.Context) ([]entity.UserDetail, error)
	//User View
	ViewSchedule(ctx context.Context, userID int) ([]entity.Event, error)
	ViewPendingPayment(ctx context.Context, userID int) ([]entity.Invoice, error)
	Login(ctx context.Context, email, password string) (*entity.User, error)
	UpdateMember(ctx context.Context, userId int, user *entity.UserDetail) error
}

type UserUsecase interface {
	Create(
		email, password, firstName, lastName, address string,
		memberTierId int,
	) error
	MemberList() ([]entity.UserDetail, error)
	//UserView
	ViewSchedule(userID int) ([]entity.Event, error)
	ViewPendingPayment(userID int) ([]entity.Invoice, error)
	Login(email, password string) (*entity.User, error)
	UpdateMember(userId int, user *entity.UserDetail) error
}

type UserHandler interface {
	Create()
	List()
	Login() (*entity.User, error)
	Update()
}
