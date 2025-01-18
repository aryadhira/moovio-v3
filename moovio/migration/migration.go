package migration

import (
	"context"
	"log"
	"moovio-v3/moovio/storages"
)

type MoovioMigration struct {
	db storages.MoovioRepo
}

func New(db storages.MoovioRepo) *MoovioMigration {
	return &MoovioMigration{
		db: db,
	}
}

func (m *MoovioMigration) getMigrateFunc(ctx context.Context) []func() {
	migrations := []func(){}
	migrations = append(migrations, func() { m.db.InitiateTable(ctx) })

	return migrations
}

func (m *MoovioMigration) Run(ctx context.Context) error {
	version, err := m.db.GetDBVersion(ctx)
	if err != nil {
		log.Println(err.Error())
		if err.Error() == `pq: relation "db_version" does not exist` {
			version = 0
		} else {
			return err
		}

	}
	log.Println("current version:", version)

	migrateFunc := m.getMigrateFunc(ctx)
	if len(migrateFunc) == version {
		log.Println("no migration needed")
	} else if version < len(migrateFunc) {
		for i := version; i < len(migrateFunc); i++ {
			log.Println("run migration version:", i+1)
			migrateFunc[i]()
			if i > 0 {
				// up version
				err := m.db.InsertDBVersion(ctx, i+1)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
