//go:build mage

package main

import (
	"fmt"
	"fwd-search-api/config"
	"fwd-search-api/migration"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Init() error {
	return sh.Run("git", "config", "core.hooksPath", "githooks")
}

func loadConfig() error {
	return config.Load()
}

func initDB() error {
	return config.InitDB()
}

func Lint() error {
	r, err := sh.Output("golangci-lint", "run", "./...", "--fix")
	fmt.Println(r)

	return err
}

func Test() error {
	r, err := sh.Output("go", "test", "-v", "./...")
	fmt.Println(r)

	return err
}

func PreCommit() error {
	mg.SerialDeps(Lint, Test)

	r, err := sh.Output("go", "mod", "tidy")
	fmt.Println(r)

	return err
}

func Changelog() error {
	r, err := sh.Output("git-chglog", "-o", "CHANGELOG.md")
	fmt.Println(r)

	return err
}

type DB mg.Namespace

func (DB) Migrate() error {
	mg.SerialDeps(loadConfig, initDB)

	return migration.Migrate()
}

func (DB) Rollback() error {
	mg.SerialDeps(loadConfig, initDB)

	return migration.RollbackLast()
}

func (DB) RollbackTo(version string) error {
	mg.SerialDeps(loadConfig, initDB)

	return migration.RollbackTo(version)
}
