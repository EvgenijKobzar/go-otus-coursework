package core

import (
	"database/sql"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	gorm "gorm.io/gorm"
	"log"
	"movies_online/internal/model/catalog"
	mocks "movies_online/internal/repository/mock"
	orm "movies_online/internal/repository/postgres/gorm"
	"os"
	"reflect"
	"testing"
)

var db *gorm.DB

func TestService_GetInner(t *testing.T) {

	repo := orm.NewRepository[*catalog.Serial](db)

	service := New[*catalog.Serial](repo)
	entity := &catalog.Serial{
		Title: "Breaking Bad [test-1]",
	}
	want, _ := service.AddInner(&entity)
	id := (*want).GetId()

	tests := []struct {
		name      string
		args      int
		want      *catalog.Serial
		wantError string
	}{
		{
			name:      "getById - success",
			args:      id,
			want:      *want,
			wantError: "",
		},
		{
			name:      "not found - success",
			args:      10000,
			want:      nil,
			wantError: "entity not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := service.GetInner(tt.args)

			if err != nil {
				if err.Error() != tt.wantError {
					t.Errorf("GetInner() error = %v, wantError = %v", err, tt.wantError)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetInner(%d) got = %v, want = %v", id, got, tt.want)
				}

				err = service.DeleteInner(tt.args)
				if err != nil {
					t.Errorf("DeleteInner() error = %v", err)
				}
			}
		})
	}
}

func TestService_GetListInner(t *testing.T) {
	t.Run("Count", func(t *testing.T) {
		repo := orm.NewRepository[*catalog.Serial](db)
		service := New[*catalog.Serial](repo)

		items, _ := service.GetListInner(make(map[string]string), make(map[string]string))
		assert.Equal(t, len(items), repo.Count(), "getList() %d; repo.Count() %d", len(items), repo.Count())

		var err error
		assert.NoError(t, err)
	})
}

func TestService_DeleteInner(t *testing.T) {
	repo := orm.NewRepository[*catalog.Serial](db)
	service := New[*catalog.Serial](repo)
	entity := &catalog.Serial{
		Title: "Breaking Bad [test-1]",
	}
	want, _ := service.AddInner(&entity)
	id := (*want).GetId()

	tests := []struct {
		name      string
		args      int
		wantError string
	}{
		{
			name:      "daleted - success",
			args:      id,
			wantError: "",
		},
		{
			name:      "not found - success",
			args:      10000,
			wantError: "entity not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := service.DeleteInner(tt.args)

			if (err != nil) && err.Error() != tt.wantError {
				t.Errorf("Delete() error = %v, wantError %v", err, tt.wantError)
				return
			}
		})
	}
}

func TestService_AddInner(t *testing.T) {
	repo := orm.NewRepository[*catalog.Serial](db)
	tests := []struct {
		name string
		args *catalog.Serial
	}{
		{
			name: "success",
			args: &catalog.Serial{
				Title: "Breaking Bad [test-1]",
			},
		},
		{
			name: "success",
			args: &catalog.Serial{
				Title: "Breaking Bad [test-2]",
			},
		},
		{
			name: "success",
			args: &catalog.Serial{
				Title: "Breaking Bad [test-3]",
			},
		},
		{
			name: "success",
			args: &catalog.Serial{
				Title: "Breaking Bad [test-4]",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			service := New[*catalog.Serial](repo)
			item, err := service.AddInner(&tt.args)

			if err != nil {
				t.Errorf("AddInner() error = %v", err)
			}

			if (*item).GetId() == 0 {
				t.Errorf("Entity not created")
			}

			err = service.DeleteInner((*item).GetId())
			if err != nil {
				t.Errorf("DeleteInner() error = %v", err)
			}
		})
	}
}

func TestService_UpdateInner(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mocks.NewMockIRepository[*catalog.Serial](ctrl)
	service := New[*catalog.Serial](repoMock)

	tests := []struct {
		name    string
		mock    func()
		fields  map[string]any
		want    *catalog.Serial
		wantErr bool
	}{
		{
			name: "success",
			mock: func() {
				repoMock.EXPECT().Save(gomock.Any()).Return(nil)
				repoMock.EXPECT().GetById(1).Return(
					&catalog.Serial{
						Id:               1,
						Title:            "Breaking Bad",
						FileId:           0,
						Description:      "",
						Rating:           9.5,
						Duration:         0,
						Sort:             0,
						ProductionPeriod: "2008-2013",
						Quality:          "",
					}, nil)
			},
			fields: map[string]any{
				"rating": 9.7,
			},
			want: &catalog.Serial{
				Id:               1,
				Title:            "Breaking Bad",
				FileId:           0,
				Description:      "",
				Rating:           9.7,
				Duration:         0,
				Sort:             0,
				ProductionPeriod: "2008-2013",
				Quality:          "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, _ := service.UpdateInner(1, tt.fields)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateInner() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func loadEnv() {
	fmt.Println("env")
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	loadEnv()

	var err error
	db, err = orm.DbConnect()
	if err != nil {
		log.Fatal(err)
	}
}

func shutdown() {
	var sqlDB *sql.DB
	var err error

	sqlDB, err = db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
