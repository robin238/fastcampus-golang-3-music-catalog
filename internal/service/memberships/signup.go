package memberships

import (
	"errors"

	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s service) SignUp(request memberships.SignUpRequest) error {

	existingUser, err := s.repository.GetUser(request.Email , request.Username, 0)
	if err !=nil {
		log.Error().Err(err).Msg("user already exist")
		return err
	}

	if existingUser != nil {
		return errors.New("email or username exists")
	}
	
	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	
	if err !=nil {
		log.Error().Err(err).Msg("Error hash password")
		return err
	}

	model :=memberships.User{
		Email: request.Email,
		Username: request.Username,
		Password: string(pass),
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}

	return s.repository.CreateUser(model)

}