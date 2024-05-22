package migrations

import "github.com/pbentes/80_20/src/db"

/*
	This package assumes the migrations were all successful.
	If it wasn't git gud. The skill issues are yours and yours alone.
	Learn SQL.
*/

type Migration struct {
	Name string
	Up   string
	Down string
}

func (migration Migration) Run() error {
	db := db.Database

	var count int
	err := db.QueryRow(db.Rebind("SELECT COUNT(*) FROM migrations WHERE name = ?"), migration.Name).Scan(&count)
	if err != nil {
		return err
	} else if count > 0 {
		return nil
	}

	_, err = db.Exec(migration.Up)
	if err != nil {
		return err
	}

	db.Exec(db.Rebind("INSERT INTO migrations (name) VALUES (?)"), migration.Name)

	return nil
}

func (migration Migration) Revert() error {
	db := db.Database

	_, err := db.Exec(migration.Down)
	if err != nil {
		return err
	}

	return nil
}

var Migrations []Migration = make([]Migration, 0)

func AddMigration(migration Migration) {
	Migrations = append(Migrations, migration)
}

func RunMigrations() error {

	migrations_table := `
		CREATE TABLE IF NOT EXISTS migrations (
			id INTEGER PRIMARY KEY,
			name TEXT UNIQUE,
			installed_on DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := db.Database.Exec(migrations_table)
	if err != nil {
		return err
	}

	for _, mig := range Migrations {
		err := mig.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func RevertMigrations() error {
	for i := range Migrations {
		err := Migrations[len(Migrations)-1-i].Revert()
		if err != nil {
			return err
		}
	}
	return nil
}

func RevertLastMigration() error {
	return Migrations[len(Migrations)-1].Revert()
}
