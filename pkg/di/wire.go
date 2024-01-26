//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ashiqsabith123/notification-svc/pkg/config"
	"github.com/ashiqsabith123/notification-svc/pkg/db"
	"github.com/ashiqsabith123/notification-svc/pkg/repository"
	"github.com/ashiqsabith123/notification-svc/pkg/service"
	"github.com/ashiqsabith123/notification-svc/pkg/usecase"
	"github.com/google/wire"
)

func IntializeService(config config.Config) service.UserService {

	wire.Build(
		db.ConnectToDatabase,
		repository.NewUserRepo,
		usecase.NewUserUsecase,
		service.NewUserService,
	)

	return service.UserService{}

}
