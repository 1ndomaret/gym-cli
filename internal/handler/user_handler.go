package handler

import "gym-cli/internal/domain"

type userHandler struct {
	uc domain.UserUsecase
}

func NewUserHandler() domain.UserHandler {
	return &userHandler{}
}

func (h *userHandler) Create() {

}

func (h *userHandler) List() {

}
