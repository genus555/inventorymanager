# **🎒Inventory Manager**

An Inventory Manager that lets you organize stock into different categories and allows you to view or update stock amounts. 
  
## 😤Motivation  
Every game has a built in inventory or some way to view and manipulate your items. I'm currently playing a public beta on a game without an inventory feature so I decided to create the Inventory Manager as a way to keep stock of what I do and don't have. While building Inventory Manager I kept the program broad so that it can be used for a general purpose rather than specifically my game. For available commands, view the Usage sections. (For further commands view the DEV.md file) (<- WIP)

## 🤖Requirements:
- [Go](https://golang.org/) 1.21+

## Quick Start (Linux)  
1. Navigate over to the project's root directory in your terminal (e.g. `cd c/Users/your_username/inventorymanager`)
2. `go run .` from the spa directory.
3. Inventory Manager will start and the user can follow the available commands to create and populate their inventory.  
  
## :pencil::thinking:Usage:  
List of available commands. Note: Each command can also be written in short hand (provided next to the name)
***!!! Notice !!!***  
*Copy the `inventory.db` file to create a backup. Simply reinsert the file to load the database at that point in time.*  
### New (n):  
*Usage: `new [category_name]`*  
Creates a new category  
### List (ls):  
*Usage: `list`*  
List all categories (Default option but can also be called with flag -categories [-c])  
**List Entries:**  
With flag, list all entries inside currently open category  
*Usage: `list -entries (-e)`*  
### Open (o):  
*Usage: `open [category_name]`*  
Opens the category for use with other commands  
### Add (a):  
*Usage: `add [entry_name]`*  
Adds an entry to the currently open category  
### Delete (d):  
*Usage: `delete [entry_name]`*  
Deletes an entry from the currently open category  
**Delete Category:**  
With flag, deletes the category from database  
*Usage: `delete -t [category_name]`*  
### Get (g):  
*Usage: `get [entry_name]`*  
Shows the amount of entry in stock  
### Update (u):  
*Usage: `update [entry_name] [new_amount]`*  
Updates the amount of entry in stock  
### Plus (p):  
*Usage: `plus [entry_name]`*  
Adds 1 to the amount in stock of entry  
### Minus (m):  
*Usage: `minus [entry_name]`*  
Subtracts 1 to the amount in stock of entry  
### Help (h):  
*Usage: `help`*  
Provides a list of available commands and how to use them  
### Quit (q):  
*Usage: `quit`*  
Stops the program  

## Contributing  
### Submit a pull request  
If you'd like to contribute, fork the repository and open a pull request to the `main` branch.
