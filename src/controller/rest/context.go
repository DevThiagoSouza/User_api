package controller

import (
	"modulo/src/domain"
	"modulo/src/exeption/validation"
	"modulo/src/logger"
	"modulo/src/model/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UseDomainInterface domain.UserDomainInterface
)

func UserId(c *gin.Context) {

}

func UserEmail(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	logger.Info("Inicio do Create User controller", zap.String("journey", "createUser"))
	var userRequest = request.UserRequest{}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao fazer a validacao", err, zap.String("journey", "createUser"))
		errR := validation.ValidationError(err)

		c.JSON(errR.Code, errR)
		return
	}

	domain := domain.NewUserDomain(
		userRequest.Email,
		userRequest.Nome,
		userRequest.Password,
		userRequest.SobreNome,
	)
	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Usuario criado com sucesso", zap.String("journey", "createUser"))
	c.String(http.StatusOK, "")
}

func UpdateUser(c *gin.Context) {}

func DeleteUser(c *gin.Context) {}
