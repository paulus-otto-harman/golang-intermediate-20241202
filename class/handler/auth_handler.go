package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/class/database"
	"project/class/domain"
	"project/class/service"
)

type AuthController struct {
	service service.AuthService
	logger  *zap.Logger
	cacher  database.Cacher
}

func NewUserHandler(service service.AuthService, logger *zap.Logger, rdb database.Cacher) *AuthController {
	return &AuthController{service: service, logger: logger, cacher: rdb}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		responseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	isAuthenticated, err := ctrl.service.Login(user)
	if err != nil {
		BadResponse(c, "server error", http.StatusInternalServerError)
		return
	}

	if !isAuthenticated {
		BadResponse(c, "authentication failed", http.StatusUnauthorized)
		return
	}

	token := "2323232"
	IDKEY := user.Username

	err = ctrl.cacher.Set(IDKEY, token)
	if err != nil {
		BadResponse(c, "server error", http.StatusInternalServerError)
	}

	SuccessResponseWithData(c, "login successfully", http.StatusOK, gin.H{"token": token, "id_key": IDKEY})
}
