package main

import (
	"GoComments/migration"
	"flag"
	"log"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "generate a migration to db")
	flag.Parse()
	if migrate == "yes" {
		log.Println("start the migration")
		migration.Migrate()
		log.Println("migration finished")
	}
}
