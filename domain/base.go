package domain

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type UsuarioId struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at"`
}

func (base *UsuarioId) AposCriacaoNovoUsuario(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error durante ao criar novo Usuario: %v", err)
	}
	return nil
}
