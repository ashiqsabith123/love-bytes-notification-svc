package usecase

import (
	repo "github.com/ashiqsabith123/notification-svc/pkg/repository/interface"
	usecase "github.com/ashiqsabith123/notification-svc/pkg/usecase/interface"
)

type UserUsecase struct {
	UserRepo repo.UserRepo
}

func NewUserUsecase(repo repo.UserRepo) usecase.UserUsecase {
	return &UserUsecase{UserRepo: repo}
}
