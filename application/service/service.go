package service

import (
	"github.com/brunotitata/go-api/application/repository"
	"github.com/brunotitata/go-api/domain"
	"github.com/brunotitata/go-api/framework/db"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func Salvar(c *gin.Context) uuid.UUID {
	conexaoDB := db.ConexaoDB()
	repositoryDB := repository.UserRepositoryDB{
		Db: conexaoDB,
	}

	usuario, err := domain.NovoUsuario(&domain.Usuario{})

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	usuarioCriado, err := repositoryDB.Insert(usuario)

	if err != nil {
		log.Fatalf("Error ao criar usuario: %v", err)
	}

	return usuarioCriado.ID
}

func BuscarUsuarioPorId(c *gin.Context) domain.Usuario {
	conexaoDB := db.ConexaoDB()
	repositoryDB := repository.UserRepositoryDB{
		Db: conexaoDB,
	}

	id := c.Query("id")
	usuario := repositoryDB.BuscarUsuarioPorId(id)

	return usuario
}

func BuscarTodosOsFuncionarios() []domain.Usuario {
	conexaoDB := db.ConexaoDB()
	repositoryDB := repository.UserRepositoryDB{
		Db: conexaoDB,
	}
	return repositoryDB.BuscarTodosFuncionarios()
}
