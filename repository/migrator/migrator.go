package migrator

import (
	"database/sql"
	"flexy/repository/sqlite"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	dbConfig   sqlite.Config
	migrations *migrate.FileMigrationSource
}

// TODO - set migration table name
// TODO - add limit to Up and Down method

func New(dbConfig sqlite.Config) Migrator {
	// OR: Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: "./repository/sqlite/migrations",
	}

	return Migrator{dbConfig: dbConfig, dialect: "sqlite3", migrations: migrations}
}

func (m Migrator) Up() {
	db, err := sql.Open(m.dialect, m.dbConfig.FilePath)
	if err != nil {
		panic(fmt.Errorf("can't open sqlite db: %v", err))
	}
	defer db.Close()

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't apply migrations: %v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {
	db, err := sql.Open(m.dialect, m.dbConfig.FilePath)
	if err != nil {
		panic(fmt.Errorf("can't open sqlite db: %v", err))
	}
	defer db.Close()

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback migrations: %v", err))
	}
	fmt.Printf("Rollbacked %d migrations!\n", n)
}

func (m Migrator) Status() {
	// TODO - add status
}
