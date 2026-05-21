package main

import (
	"fmt"
	"log"
	"database/sql"

	_			"modernc.org/sqlite"
	database	"github.com/genus555/inventorymanager/internal/database"
	cli			"github.com/genus555/inventorymanager/internal/cli"
)

func main() {
	fmt.Println("Welcome to Inventory Manager")
	db_file, err := sql.Open("sqlite", "./inventory.db")
	if err != nil {
		log.Fatalf("Problem creating database: %v", err)
	}
	defer db_file.Close()

	db := database.NewDB(db_file)
	fmt.Println(db.TableName)

	cli.PrintCommands()
	
	for {
		inputs := cli.GetInput()
		if len(inputs) == 0 {
			continue
		}
		switch inputs[0] {
		case "dev":
			err := HandleDevOptions(db, inputs)
			if err != nil {fmt.Println(err)}
		case "list":
			fallthrough
		case "ls":
			err := HandleList(db, inputs)
			if err != nil {fmt.Println(err)}
		case "new":
			fallthrough
		case "n":
			err := HandleCreateTable(db, inputs)
			if err != nil {fmt.Println(err)}
		case "open":
			fallthrough
		case "o":
			err := HandleOpen(db, inputs)
			if err != nil {fmt.Println(err)}
		case "add":
			fallthrough
		case "a":
			err := HandleCreateEntry(db, inputs)
			if err != nil {fmt.Println(err)}
		case "delete":
			fallthrough
		case "d":
			err := HandleDeleteEntry(db, inputs)
			if err != nil {fmt.Println(err)}
		case "get":
			fallthrough
		case "g":
			err := HandleGetEntry(db, inputs)
			if err != nil {fmt.Println(err)}
		case "update":
			fallthrough
		case "u":
			err := HandleUpdateEntry(db, inputs)
			if err != nil {fmt.Println(err)}
		case "plus":
			fallthrough
		case "p":
			fallthrough
		case "minus":
			fallthrough
		case "m":
			err := HandlePlusMinus(db, inputs)
			if err != nil {fmt.Println(err)}
		case "restock":
			fallthrough
		case "re":
			err := HandleCheckRestock(db, inputs)
			if err != nil {fmt.Println(err)}
		case "h":
			fallthrough
		case "help":
			cli.PrintCommands()
		case "q":
			fallthrough
		case "quit":
			fmt.Println("Closing Inventory Manager")
			return
		default:
			fmt.Printf("\"%s\" is not a valid command\n", inputs[0])
		}
	}
}