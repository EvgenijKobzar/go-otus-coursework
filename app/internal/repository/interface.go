package repository

import "movies_online/internal/model/catalog"

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mocks

type IRepository[T catalog.HasId] interface {
	//GetAll() ([]T, error)
	GetAll(filter map[string]string, order map[string]string) ([]T, error)
	GetById(id int) (T, error)
	Save(entity T) error
	Delete(id int) error
	Count() int
}
