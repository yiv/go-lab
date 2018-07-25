package main

import (
	"fmt"
	"log"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Account is our model, which corresponds to the "accounts" database table.
type AccountInfo struct {
	ID      int `gorm:"primary_key"`
	Balance int
}

func main() {
	// Connect to the "bank" database as the "maxroach" user.
	const addr = "postgresql://root@192.168.1.51:26257/bank?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Automatically create the "accounts" table based on the Account model.
	db.AutoMigrate(&AccountInfo{})

	// Insert two rows into the "accounts" table.
	db.Create(&AccountInfo{ID: 1, Balance: 1000})
	db.Create(&AccountInfo{ID: 2, Balance: 250})

	// Print out the balances.
	var accounts []AccountInfo
	db.Find(&accounts)
	fmt.Println("Initial balances:")
	for _, AccountInfo := range accounts {
		fmt.Printf("%d %d\n", AccountInfo.ID, AccountInfo.Balance)
	}
}
