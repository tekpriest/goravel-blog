package providers

import (
	"goravel/app/models"
	"goravel/database/seeders"

	"github.com/goravel/framework/contracts/database/seeder"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/database/gorm"
	"github.com/goravel/framework/facades"
)

type DatabaseServiceProvider struct{}

func (receiver *DatabaseServiceProvider) Register(app foundation.Application) {}

func (receiver *DatabaseServiceProvider) Boot(app foundation.Application) {
	db := facades.Orm().Query().(*gorm.QueryImpl).Instance()
	db.AutoMigrate(&models.User{}, &models.Post{})
	facades.Seeder().Register([]seeder.Seeder{
		&seeders.DatabaseSeeder{},
	})
}
