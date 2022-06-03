package store

import (
	"fmt"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"net/http"
	"sf6/internal/sqlScript"
)

func (s *Store) migrateDatabase() error {
	logrus.Infof("Updating")

	// Retrieve update scripts to apply
	migrate.SetTable("migration")
	migrationSource := migrate.HttpFileSystemMigrationSource{
		FileSystem: http.FS(sqlScript.Fs),
	}
	updateScripts, err := migrationSource.FindMigrations()
	if err != nil {
		return fmt.Errorf("can't find update scripts: %w", err)
	}

	if len(updateScripts) == 0 {
		logrus.Info("no update scripts to apply")
	}

	// Apply update scripts
	n, err := migrate.Exec(s.db.DB, "postgres", migrationSource, migrate.Up)
	if err != nil {
		return fmt.Errorf("can't run migrations: %w", err)
	}

	logrus.Infof("%d update scripts applied", n)

	return nil
}
