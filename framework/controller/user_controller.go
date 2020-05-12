package controller

import (
	"github.com/brunotitata/go-api/application/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CriarNovoUsuario() gin.HandlerFunc {
	return func(c *gin.Context) {
		idUsuario := service.Salvar(c)
		c.JSON(http.StatusCreated, gin.H{
			"id": idUsuario,
		})
	}
}

func BuscarUsuarioPorId() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarios := service.BuscarUsuarioPorId(c)
		c.JSON(http.StatusOK, gin.H{
			"usuarios": usuarios,
		})
	}
}

func BuscarTodosUsuarios() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarios := service.BuscarTodosOsFuncionarios()
		c.JSON(http.StatusOK, usuarios)
	}
}
