package migrate

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"github.com/mattes/migrate/source/go-bindata"
)

// Down migrates all the way down
func Down(databaseUrl string) error {
	assetSource := bindata.Resource(
		AssetNames(),
		func(name string) ([]byte, error) {
			return Asset(name)
		},
	)
	sourceDriver, err := bindata.WithInstance(assetSource)
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return err
	}
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("go-bindata", sourceDriver, "postgres", dbDriver)
	if err != nil {
		return err
	}
	return m.Down()
}

// Up migrates all the way up
func Up(databaseUrl string) error {
	assetSource := bindata.Resource(
		AssetNames(),
		func(name string) ([]byte, error) {
			return Asset(name)
		},
	)
	sourceDriver, err := bindata.WithInstance(assetSource)
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return err
	}
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("go-bindata", sourceDriver, "postgres", dbDriver)
	if err != nil {
		return err
	}
	return m.Up()
}
