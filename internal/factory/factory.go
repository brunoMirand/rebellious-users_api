package factory

import (
	"github.com/brunoMirand/rebellious-users_api/internal/usecase"
  "github.com/brunoMirand/rebellious-users_api/internal/repository"
)

type UserFactory struct {
  createUserUseCase *usecase.CreateUserUseCase
  findAllUsersUseCase *usecase.FindAllUsersUseCase
}

func NewUserFactory() *UserFactory {
  repo := repository.NewInMemoryUsersRepository()
  return &UserFactory{
    createUserUseCase: usecase.NewCreateUserUseCase(repo),
    findAllUsersUseCase: usecase.NewFindAllUsersUseCase(repo),
  }
}

func (f *UserFactory) CreateUserUseCase() *usecase.CreateUserUseCase {
  return f.createUserUseCase
}

func (f *UserFactory) FindAllUsersUseCase() *usecase.FindAllUsersUseCase {
  return f.findAllUsersUseCase
}
