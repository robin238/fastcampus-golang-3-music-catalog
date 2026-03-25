package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/models/memberships"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Login(c *gin.Context) {
	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("error bind json")
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	accessToken, err := h.service.Login(request)
	if err != nil {
		log.Error().Err(err).Msg("error login")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})

}
