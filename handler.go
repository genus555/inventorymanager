package main

import (
	"fmt"
	"strconv"

	database	"github.com/genus555/inventorymanager/internal/database"
)

func strToInt(num string) (int, error) {
	convNum, err := strconv.Atoi(num)
	if err != nil {return 0, err}
	return convNum, nil
}

func HandleCreateTable(db *database.DB, inputs []string) error {
	if len(inputs) < 2 {
		return fmt.Errorf("Incorrect usage\nUsage: new (or 'n') [category_name]")
	}

	err := db.CreateTable(inputs)
	if err != nil {return err}
	return nil
}

func HandleList(db *database.DB, inputs []string) error {
	if len(inputs) == 1 || (len(inputs) >= 2 && (inputs[1] == "-c" || inputs[1] == "-categories")) {
		err := db.ListTables()
		if err != nil {return err}
		return nil
	}
	if inputs[1] == "-e" || inputs[1] == "-entries" {
		if len(inputs) == 2 {
			if db.TableName == "" {return fmt.Errorf("No category is currently open.")}
			err := db.ListEntries(db.TableName)
			if err != nil {return err}
		} else {
			tableName, err := db.CheckTable(inputs[2])
			if err != nil {return err}
			err = db.ListEntries(tableName)
			if err != nil {return err}
		}
	} else {return fmt.Errorf("Unknown flag")}

	return nil
}

func HandleOpen(db *database.DB, inputs []string) error {
	if len(inputs) < 2 {return fmt.Errorf("Incorrect usage. Usage: open [category_name]")}
	table, err := db.CheckTable(inputs[1])
	if err != nil {return err}
	db.TableName = table
	fmt.Printf("Category \"%s\" has been opened.\n", db.TableName)
	return nil
}

func HandleCreateEntry(db *database.DB, inputs []string) error {
	if db.TableName == "" {return fmt.Errorf("No category is currently open.")}
	if len(inputs) < 2 {
		return fmt.Errorf("Incorrect usage. Usage: add [entry_name]")
	}
	err := db.AddEntry(inputs[1])
	if err != nil {return err}
	return nil
}

func HandleDeleteEntry(db *database.DB, inputs []string) error {
	if len(inputs) < 2 {
		return fmt.Errorf("Incorrect usage. Usage: delete [entry_name]")
	}
	if len(inputs) >= 3 {
		if inputs[1] == "-t" {
			err := db.DeleteTable(inputs[2])
			if err != nil {return err}
			return nil
		}
	}

	if db.TableName == "" {return fmt.Errorf("No category is currently open.")}
	err := db.DeleteEntry(inputs[1])
	if err != nil {return err}
	return nil
}

func HandleUpdateEntry(db *database.DB, inputs []string) error {
	if db.TableName == "" {return fmt.Errorf("No category is currently open.")}
	if len(inputs) < 3 {
		return fmt.Errorf("Incorrect usage. Usage: update [entry_name] [amount]")
	}
	num, err := strToInt(inputs[2])
	if err != nil {return err}
	err = db.UpdateEntry(inputs[1], num)
	if err != nil {return err}
	return nil
}

func HandlePlusMinus(db *database.DB, inputs []string) error {
	if db.TableName == "" {return fmt.Errorf("No category is currently open.")}
	if len(inputs) < 2 {
		return fmt.Errorf("Incorrect usage. Usage: plus/minus [entry_name]")
	}
	if inputs[0] == "plus" || inputs[0] == "p" {
		db.PlusMinus(inputs[1], database.PLUS)
		return nil
	}
	if inputs[0] == "minus" || inputs[0] == "m" {
		db.PlusMinus(inputs[1], database.MINUS)
		return nil
	}
	return fmt.Errorf("Something wrong has happened with Plus Minus")
}

func HandleGetEntry(db *database.DB, inputs []string) error {
	if db.TableName == "" {return fmt.Errorf("No category is currently open.")}
	if len(inputs) < 2 {
		return fmt.Errorf("Incorrect usage. Usage: get [entry_name]")
	}

	amount, err := db.GetEntry(inputs[1])
	if err != nil {return err}

	fmt.Printf("%s \"%s\": %d\n", db.TableName, inputs[1], amount)
	return nil
}

func HandleCheckRestock(db *database.DB, inputs []string) error {
	if db.TableName == "" {return fmt.Errorf("No category is currently open.")}
	if len(inputs) == 1 {
		if err := db.GetLow(); err != nil {return err}
		if err := db.GetEmpty(); err != nil {return err}
	}else if inputs[1] == "-l" || inputs[1] == "-low" {
		if err := db.GetLow(); err != nil {return err}
	} else if inputs[1] == "-e" || inputs[1] == "-empty" {
		if err := db.GetEmpty(); err != nil {return err}
	} else {
		return fmt.Errorf("Incorrect usage. Usage: restock (optional flags -empty/-low)")
	}
	return nil
}

func HandleDevOptions(db *database.DB, inputs []string) error {
	if len(inputs) < 2 {return fmt.Errorf("Missing Dev Flag")}
	switch inputs[1] {
	case "wipe":
		if len(inputs) > 2 {
			if inputs[2] == "category" {
				if len(inputs) > 3 {
					err := db.WipeTable(inputs[3])
					if err != nil {return err}
				} else {return fmt.Errorf("Missing category name")}
				return nil
			}
		}
		err := db.Wipe()
		if err != nil {return err}
	default:
		return fmt.Errorf("Incorrect Flag")
	}
	return nil
}