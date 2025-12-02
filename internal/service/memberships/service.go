package memberships

import (
	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/configs"
	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/models/memberships"
)

type repository interface {
	CreateUser(model memberships.User) error
	GetUser(email, username string, id int) (*memberships.User, error)
}

type service struct {
	cfg *configs.Config
	repository repository

}

func NewService(cfg *configs.Config, repository repository) *service {
	return &service{
		cfg:cfg,
		repository: repository,
	}
}