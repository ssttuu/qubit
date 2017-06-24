package migrationutils

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"os"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"github.com/mattes/migrate/source/go-bindata"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	migrateTask      = kingpin.Command("migrate", "Direction in which to migrate the DB")
	migrateDirection = migrateTask.Arg("direction", "Direction in which to migrate the DB").Enum("down", "up")
)

func migrateDB(direction string, assetNames []string, afn bindata.AssetFunc) error {
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		return errors.New(`You need to set the environment variable "POSTGRES_URL"`)
	}
	assetSource := bindata.Resource(assetNames, afn)
	sourceDriver, err := bindata.WithInstance(assetSource)
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", postgresURL)
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
	if direction == "up" {
		return m.Up()
	} else if direction == "down" {
		return m.Down()
	}
	return errors.New("Invalid migration direction: must be up or down")
}

// Init loads and runs migrations as appropriate
func Init(assetNames []string, afn bindata.AssetFunc) {
	switch kingpin.Parse() {
	case "migrate":
		err := migrateDB(*migrateDirection, assetNames, afn)
		if err != nil {
			log.Fatal(err)
		}
		break
	}
}
