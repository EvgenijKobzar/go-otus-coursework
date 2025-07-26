package gorm

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"movies_online/internal/model/catalog"
	"os"
	"strings"
	"time"
)

func NewRepository[T catalog.HasId](db *gorm.DB) *Repository[T] {
	return &Repository[T]{
		Db: db,
	}
}

func (r *Repository[T]) Save(entity T) error {
	var err error

	if entity.GetId() == 0 {
		err = r.add(entity)
	} else {
		err = r.update(entity)
	}
	return err
}

func (r *Repository[T]) add(entity T) error {
	result := r.Db.Create(&entity)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) update(entity T) error {
	var err error

	result := r.Db.Save(&entity)
	if err = result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("entity not found")
		}
	}
	return err
}

func (r *Repository[T]) Delete(id int) error {
	var err error
	var entity T

	result := r.Db.Delete(&entity, id)
	if err = result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("entity not found")
		}
	}
	return err
}

func (r *Repository[T]) GetAll(filterFields map[string]string, orderFields map[string]string) ([]T, error) {
	var items []T
	var err error

	fmt.Print(filterFields)
	if val, ok := filterFields["%description"]; ok {
		result := r.Db.Where("description LIKE ?", "%"+val+"%").Find(&items)
		if result.Error != nil {
			err = result.Error
		}
	} else {
		filter, err := r.makeFilter(filterFields)
		if err != nil {
			return nil, err
		}

		result := r.Db.Order(makeSort(orderFields)).Find(&items, filter)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	return items, err
}

func (r *Repository[T]) makeFilter(fields map[string]string) (T, error) {
	var err error
	filter := new(T)

	config := &mapstructure.DecoderConfig{
		Result:           filter,
		WeaklyTypedInput: true,
	}

	var d *mapstructure.Decoder
	d, err = mapstructure.NewDecoder(config)
	if err == nil {
		err = d.Decode(fields)
	}

	return *filter, err
}

func makeSort(fields map[string]string) string {
	list := make([]string, 0)
	var builder strings.Builder
	for k, v := range fields {
		builder.WriteString(k)
		builder.WriteString(" ")
		builder.WriteString(v)
		list = append(list, builder.String())
		builder.Reset()
	}

	separator := ", "
	return strings.Join(list, separator)
}

func (r *Repository[T]) GetById(id int) (T, error) {
	var entity T
	var err error

	err = r.Db.First(&entity, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("entity not found")
	}
	return entity, err
}

func (r *Repository[T]) Count() int {
	items, _ := r.GetAll(make(map[string]string), make(map[string]string))
	return len(items)
}

func DbConnect() (*gorm.DB, error) {
	if db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_DB_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: os.Getenv("POSTGRES_DB_TABLE_PREFIX"),
		},
	}); err == nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(50)               // Рекомендуется: (max_connections / worker_processes) * 0.8
		sqlDB.SetMaxIdleConns(15)               // ~25-50% от MaxOpenConns
		sqlDB.SetConnMaxLifetime(1 * time.Hour) // Зависит от настроек БД
		sqlDB.SetConnMaxIdleTime(15 * time.Minute)
		return db, nil
	} else {
		return nil, err
	}
}
