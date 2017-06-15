package main

import (
	"errors"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/stupschwartz/qubit/applications/api/tasks/migrate"
)

var (
	migrateTask      = kingpin.Command("migrate", "Direction in which to migrate the DB")
	migrateDirection = migrateTask.Arg("direction", "Direction in which to migrate the DB").Enum("down", "up")
)

func migrateDatabase(direction string) error {
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		return errors.New(`You need to set the environment variable "POSTGRES_URL"`)
	}
	if direction == "up" {
		return migrate.Up(postgresURL)
	} else if direction == "down" {
		return migrate.Down(postgresURL)
	}
	return errors.New("Invalid migration direction: must be up or down")
}

func main() {
	switch kingpin.Parse() {
	case "migrate":
		err := migrateDatabase(*migrateDirection)
		if err != nil {
			log.Fatal(err)
		}
		break
	}
}
