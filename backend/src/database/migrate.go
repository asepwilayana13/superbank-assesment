package database

import (
	"fmt"
	"log"
	"myapp-go/src/models"
)

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.User{}, 
		&models.Session{},
		&models.Customer{},
		&models.BankAccount{},
		&models.Pocket{},
		&models.TermDeposit{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database migrated successfully!")
}