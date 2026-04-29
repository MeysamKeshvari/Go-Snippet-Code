package main

import (
	"flag"
	"fmt"
	"os"
)
type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
}

func main() {

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "List all the books")
	getId := getCmd.String("id", "", "Get book by id")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addId := addCmd.String("id", "", "Book id")
	addTitle := addCmd.String("title", "", "Book title")
	addAuthor := addCmd.String("author", "", "Book author")
	addPrice := addCmd.String("price", "", "Book price")
	
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.String("id", "", "Book id")
	updateTitle := updateCmd.String("title", "", "Book title")
	updateAuthor := updateCmd.String("author", "", "Book author")
	updatePrice := updateCmd.String("price", "", "Book price")
	
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.String("id", "", "Delete book by id")

	if len(os.Args) < 2 {
		fmt.Println("Expected get, add, update, delete commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		handleGetBooks(getCmd, getAll, getId)
	case "add":
		handleAddBook(addCmd, addId, addTitle, addAuthor, addPrice,true)
	case "update":
		handleAddBook(updateCmd, updateId, updateTitle, updateAuthor, updatePrice, false)
	case "delete":
		handleDeleteBook(deleteCmd, deleteId)
	default:
		fmt.Println("Please provide get, update, update, delete commands")
		os.Exit(1)
	}
}
