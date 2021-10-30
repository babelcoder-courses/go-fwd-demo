package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1635587856CreateResourcesTable() *gormigrate.Migration {
	type resource struct {
		gorm.Model
		Title string
	}

	dst := &resource{}

	return &gormigrate.Migration{
		ID: "1635587856",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(dst)
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(dst)
		},
	}
}
