package gorm

import (
	"gorm.io/gorm"
	"movies_online/internal/model/catalog"
)

type Repository[T catalog.HasId] struct {
	Db *gorm.DB
}
