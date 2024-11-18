package migration

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

var DbConnection *sql.DB

func DBMigrate(dbParam *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "databases/migration/sql_migrations", // folder tempat file migrasi berada
	}

	// Menjalankan migrasi
	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		// Jika terjadi kesalahan, print error dan panik
		fmt.Println("Error during migration:")
		panic(errs)
	}

	// Jika berhasil, tampilkan jumlah migrasi yang diterapkan
	fmt.Println("Migration success, applied", n, "migrations!")

	// Set koneksi database global
	DbConnection = dbParam
}
