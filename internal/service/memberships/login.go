package memberships

import (
	"errors"

	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/models/memberships"
	"github.com/robin238/fastcampus-golang-3-music-catalog/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(request memberships.LoginRequest) (string, error) {
	userDetail, err := s.repository.GetUser(request.Email, "", uint(0))
	if err != nil {
		log.Error().Err(err).Msg("error get user from database")
		return "", err
	}

	if userDetail == nil {
		return "", errors.New("email not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password))
	if err != nil {
		log.Error().Err(err).Msg("error compare password")
		return "", errors.New("email and password not match")
	}

	accessToken, err := jwt.CreateToken(int64(userDetail.ID), userDetail.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create JWT token")
		return "", err
	}

	return accessToken, nil
}
