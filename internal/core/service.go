package core

import (
	"movies_online/internal/model/catalog"
	"movies_online/internal/repository"
	"movies_online/pkg/lib/mapstructure"
)

type Service[T catalog.HasId] struct {
	repo repository.IRepository[T]
}

func New[T catalog.HasId](repo repository.IRepository[T]) *Service[T] {
	return &Service[T]{repo: repo}
}

func (service *Service[T]) GetInner(id int) (T, error) {
	return service.repo.GetById(id)
}

func (service *Service[T]) GetListInner(filter map[string]string, order map[string]string) ([]T, error) {
	items, err := service.repo.GetAll(filter, order)
	return items, err
}

func (service *Service[T]) AddInner(binding *T) (*T, error) {
	var err error
	if err = service.repo.Save(*binding); err == nil {
		return binding, nil
	}
	return nil, err
}

func (service *Service[T]) UpdateInner(id int, inputFields map[string]any) (T, error) {
	var err error
	var entity T

	entity, err = service.repo.GetById(id)

	if err == nil {
		bindings := new(T)
		err = mapstructure.MapToStruct(inputFields, bindings)
		if err == nil {
			if err = entityAssign[T](entity, *bindings, inputFields); err == nil {
				err = service.repo.Save(entity)
			}
		}
	}
	return entity, err
}

func (service *Service[T]) DeleteInner(id int) error {
	var err error

	_, err = service.repo.GetById(id)

	if err == nil {
		err = service.repo.Delete(id)
	}
	return err
}

func entityAssign[T catalog.HasId](entity T, bindings T, allowedFields map[string]any) error {
	var err error
	var srcMap map[string]any
	var distMap map[string]any

	if srcMap, err = mapstructure.StructToMap(entity); err == nil {
		if distMap, err = mapstructure.StructToMap(bindings); err == nil {
			assignMap := assign(srcMap, distMap, allowedFields)
			err = mapstructure.MapToStruct(assignMap, entity)
		}
	}
	return err
}

func assign(src map[string]any, dist map[string]any, allowed map[string]any) map[string]any {
	result := make(map[string]any)
	for name, oldValue := range src {
		if _, isAllow := allowed[name]; isAllow {
			if newValue, exists := dist[name]; exists {
				result[name] = newValue
			}
		} else {
			result[name] = oldValue
		}
	}
	return result
}
