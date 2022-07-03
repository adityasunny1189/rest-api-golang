package main

import (
	"fmt"

	"github.com/adityasunny1189/rest-api-golang/internal/db"
)

func Run() error {
	fmt.Printf("running out application\n")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Printf("Failed to connect to database\n")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Printf("failed to migrate database\n")
		return err
	}
	fmt.Printf("successfully connected to database\n")
	return nil
}

func main() {
	fmt.Printf("Go REST API\n")

	if err := Run(); err != nil {
		fmt.Printf("error in running application\n")
	}
}
