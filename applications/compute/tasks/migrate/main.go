package main

import (
	"github.com/stupschwartz/qubit/applications/compute/tasks/migrate/migrations"
	"github.com/stupschwartz/qubit/applications/lib/migrationutils"
)

func main() {
	migrationutils.Init(
		migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		},
	)
}
