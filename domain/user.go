package domain

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Usuario struct {
	UsuarioId
	Nome        string    `json:"Nome" gorm:"type:varchar(255)"`
	Sobrenome   string    `json:"Sobrenome" gorm:"type:varchar(255)"`
	Email       string    `json:"Email" gorm:"type:varchar(255);unique_index"`
	Senha       string    `json:"Senha" gorm:"type:varchar(255)"`
	TokenAccess uuid.UUID `json:"TokenAcess" gorm:"type:varchar(255);unique_index"`
}

func NovoUsuario(newUser *Usuario) (*Usuario, error) {

	usuario := Usuario{
		Nome:      newUser.Nome,
		Sobrenome: newUser.Sobrenome,
		Email:     newUser.Email,
	}

	usuario.UsuarioId.ID = uuid.NewV4()
	usuario.TokenAccess = uuid.NewV4()
	password, err := bcrypt.GenerateFromPassword([]byte(newUser.Senha), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error ao gerar senha: %v", err)
		return nil, err
	}
	usuario.Senha = string(password)
	return &usuario, nil
}
