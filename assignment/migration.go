package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"
)

const VersionTable = "schema_version"

//TODO: the library doesn't look reliable enough as doesn't store checksum of migration
func Migrate(pool *pgxpool.Pool) error {
	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	migrator, err := migrate.NewMigrator(context.Background(), conn.Conn(), VersionTable)
	if err != nil {
		return err
	}
	migrator.AppendMigration(
		"Create apikeys table",
		`
CREATE TABLE "apikeys" (
"user_id" varchar(36) NOT NULL,
"exchange" varchar(36) NOT NULL,
"id" varchar(36) NOT NULL,
"api_key_preview" varchar(5) NOT NULL,
"api_key" bytea NOT NULL,
"secret" bytea NOT NULL,
PRIMARY KEY ("user_id","exchange"))`,
		"DROP TABLE apikeys")

	migrator.AppendMigration(
		"Add by id index to apikeys",
		"CREATE UNIQUE INDEX apikeys_id_idx ON apikeys (id)",
		"DROP TABLE apikeys")

	err = migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	return nil
}
