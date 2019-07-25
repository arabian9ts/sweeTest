package main

import (
	"database/sql"
	"fmt"
	"github.com/arabian9ts/sweeTest/app/builder"
	"github.com/arabian9ts/sweeTest/config"
	"github.com/arabian9ts/sweeTest/infrastructure"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	db, err := sql.Open("mysql", config.GetSettings().GetRdbUri())
	if err != nil {
		fmt.Println("failed to connect mysql")
	}

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		fmt.Println("failed to migarate")
	}
	fmt.Printf("Applied %d migrations!\n", n)

	root, err := builder.InitializeRootController()
	if err != nil {
		panic("failed to initialize controlelrs")
	}

	infrastructure.Router(root).Run()
}
