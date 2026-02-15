package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"wappi/internal/platform/config"
)

var (
	db   *sql.DB
	once sync.Once
)

// GetInstance returns the singleton database instance
func GetInstance() *sql.DB {
	once.Do(func() {
		cfg := config.GetInstance()
		var err error
		db, err = sql.Open("postgres", cfg.DatabaseURL)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}

		log.Println("Database connection established")
	})
	return db
}

func getRelativePathToMigrationsDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	absMigrationsDirPath := filepath.Join(cwd, "migrations")

	relMigrationsDirPath, err := filepath.Rel(cwd, absMigrationsDirPath)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("file://%s", relMigrationsDirPath), nil
}

// RunMigrations applies all pending database migrations
func RunMigrations() error {
	cfg := config.GetInstance()
	migrationPath, err := getRelativePathToMigrationsDirectory()
	if err != nil {
		return err
	}

	m, err := migrate.New(migrationPath, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			log.Printf("migrations: error closing source: %v", srcErr)
		}
		if dbErr != nil {
			log.Printf("migrations: error closing database: %v", dbErr)
		}
	}()

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("failed to get migration version: %w", err)
	}
	log.Printf("migrations: current version is %v (dirty: %v)", version, dirty)

	if dirty {
		log.Printf("migrations: database is in dirty state, forcing version to %v", version)
		if err := m.Force(int(version)); err != nil {
			return fmt.Errorf("failed to force migration version: %w", err)
		}
		log.Println("migrations: dirty state cleared, retrying migration")
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("migrations: no new migrations to apply")
			return nil
		}
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("migrations: database migrated successfully")

	return nil
}
