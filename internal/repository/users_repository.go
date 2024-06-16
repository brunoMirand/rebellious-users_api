package repository

import (
  "context",
  "github.com/brunoMirand/rebellious-users_api/internal/repository"
)

type UsersRepository interface {
	Save(ctx context.Context, user *models.Users) error
  FindAll(ctx context.Context) ([]*models.Users, error)
}

type InMemoryUsersRepository struct {
  users map[string]*models.Users
}

func NewInMemoryUsersRepository() UsersRepository {
  return &InMemoryUsersRepository{
    users: make(map[string]*models.Users)
  }
}

func (r *InMemoryUsersRepository) Save(ctx context.Context, user *models.Users) error {
  r.users[user.ID] = user
  return nil
}

func (r *InMemoryUsersRepository) FindAll(ctx context.Context) ([]*models.Users, error) {
  var users []*models.Users
  for _, user := range r.users {
    users = append(users, user)
  }

  return users, nil
}
