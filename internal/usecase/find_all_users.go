package usecase

import (
  "context"

	"github.com/brunoMirand/rebellious-users_api/internal/models"
  "github.com/brunoMirand/rebellious-users_api/internal/repository"
)

type FindAllUsersUseCase struct {
  Repository repository.UsersRepository
}

func NewFindAllUsersUseCase(repo repository.UsersRepository) *FindAllUsersUseCase {
  return &FindAllUsersUseCase{Repository: repo}
}

func (uc *FindAllUsersUseCase) Execute(ctx context.Context) ([]*models.Users, error) {
  return uc.Repository.FindAll(ctx)
}
