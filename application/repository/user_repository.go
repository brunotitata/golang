package repository

import (
	"github.com/brunotitata/go-api/domain"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
)

type UserRepository interface {
	Inserir(usuario *domain.Usuario) (*domain.Usuario, error)
}

type UserRepositoryDB struct {
	Db *gorm.DB
}

func (repo *UserRepositoryDB) Insert(usuario *domain.Usuario) (*domain.Usuario, error) {

	novoUsuario, err := domain.NovoUsuario(usuario)

	if err != nil {
		log.Fatalf("Error ao persistir usuario: %v", err)
		return novoUsuario, err
	}
	err = repo.Db.Create(novoUsuario).Error

	if err != nil {
		log.Fatalf("Error ao persistir usuario: %v", err)
		return novoUsuario, err
	}

	return novoUsuario, nil
}

func (repo *UserRepositoryDB) BuscarUsuarioPorId(id string) domain.Usuario {
	var usuario domain.Usuario
	userId, _ := uuid.FromString(id)
	result := repo.Db.Where("id = ?", userId).First(&usuario).Error

	if result != nil {
		log.Fatalf("Error ao buscar o usuario: %v", result)
	}

	return usuario
}

func (repo *UserRepositoryDB) BuscarTodosFuncionarios() []domain.Usuario {
	var usuarios []domain.Usuario
	err := repo.Db.Find(&usuarios).Error

	if err != nil {
		log.Fatalf("Error ao buscar todos os usuarios: %v", err)
	}
	return usuarios
}
