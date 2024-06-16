package handlers

import (
  "net/http"
  "encoding/json"

	"github.com/brunoMirand/rebellious-users_api/internal/factory"
  "github.com/brunoMirand/rebellious-users_api/internal/usecase"
)

type UserHandler struct {
  createUserUseCase *usecase.CreateUserUseCase
  findAllUsersUseCase *usecase.FindAllUsersUseCase
}

func NewUserHandler(f *factory.UserFactory) *UserHandler {
  return &UserHandler{
    createUserUseCase: f.CreateUserUseCase(),
    findAllUsersUseCase: f.FindAllUsersUseCase()
  }
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
  var input struct {
    Name string `json:"name"`
    Email string `json:"email"`
  }

  if err := json.NewDecoder(r.body).Decode(&input); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  user, err := h.createUserUseCase.Execute(r.Context(), input.Name, input.Email)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
  users, err := h.findAllUsersUseCase.Execute(r.Context())
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(users)
}
