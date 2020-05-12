package main

import (
	"github.com/brunotitata/go-api/framework/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := gin.Default()
	r.GET("/usuarios", controller.BuscarTodosUsuarios())
	r.GET("/usuario", controller.BuscarUsuarioPorId())
	r.POST("/", controller.CriarNovoUsuario())
	r.Run()
}
