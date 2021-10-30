package migration

import (
	"fwd-search-api/config"

	"github.com/go-gormigrate/gormigrate/v2"
)

func newMigrate() *gormigrate.Gormigrate {
	db := config.GetDB()

	return gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1635587856CreateResourcesTable(),
		},
	)
}

func Migrate() error {
	return newMigrate().Migrate()
}

func RollbackLast() error {
	return newMigrate().RollbackLast()
}

func RollbackTo(version string) error {
	return newMigrate().RollbackTo(version)
}
