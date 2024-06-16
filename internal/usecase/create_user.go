package usecase

import (
  "content"

  "github.com/google/uuid"
  "github.com/brunoMirand/rebellious-users_api/internal/models"
  "github.com/brunoMirand/rebellious-users_api/internal/repository"
)

type CreateUserUseCase struct {
	Repository repository.UsersRepository
}

func NewCreateUserUseCase(repo repository.UsersRepository) *CreateUserUseCase {
  return &CreateUserUseCase{Repository: repo}
}

func (uc *CreateUserUseCase) Execute(ctx content.Context, name, email string) (*models.Users, error) {
  user := &models.Users{
    ID: uuid.New().String(),
    Name: name,
    Email: email
  }

  err := uc.Repository.Save(ctx, user)
  if err != nil {
    return nil, err
  }

  return user, nil
}
