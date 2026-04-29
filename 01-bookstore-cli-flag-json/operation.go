package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"01-bookstore-cli-flag-json/models"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error Happened ", err)
		os.Exit(1)
	}
}
// get all the books
func getBooks() (books []models.Book) {
	booksBytes, err := ioutil.ReadFile("./books.json")
	checkError(err)
	err = json.Unmarshal(booksBytes, &books)
	checkError(err)
	return books
}

// save books to books.json file
func saveBooks(books []models.Book) error {

	// converting into bytes for writing into a file
	booksBytes, err := json.Marshal(books)
	checkError(err)
	err = ioutil.WriteFile("./books.json", booksBytes, 0644)
	return err
}


// get all the books logic
func handleGetBooks(getCmd *flag.FlagSet, all *bool, id *string) {
	
	getCmd.Parse(os.Args[2:])
	// checking for all or id
	if !*all && *id == "" {
		fmt.Println("subcommand --all or --id needed")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	// if for all return all books
	if *all {
		books := getBooks()
		fmt.Printf("Id \t Title \t Author \t Price \t \n")

		for _, book := range books {
			fmt.Printf("%v \t %v \t %v \t %v \t \n", book.Id, book.Title, book.Author, book.Price)
		}
	}
	// if id, return only that book
	// Throw error if book not found

	if *id != "" {
		books := getBooks()
		fmt.Printf("Id \t Title \t Author \t Price \t \n")

		var foundBook bool
		for _, book := range books {
			foundBook = true
			if *id == book.Id {
				fmt.Printf("%v \t %v \t %v \t %v \t \n", book.Id, book.Title, book.Author, book.Price)
			}
		}
		// if no book found with mentioned id throws an error
		if !foundBook {
			fmt.Println("Book not found")
			os.Exit(1)
		}
	}
}
// add or update a book logic
func handleAddBook(addCmd *flag.FlagSet, id, title, author, price *string, addNewBook bool) {
	addCmd.Parse(os.Args[2:])

	if *id == "" || *title == "" || *author == "" || *price == "" {
		fmt.Println("Please provide book id, title, author,price")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
	books := getBooks()
	var newBook models.Book
	// to check a book exist or not
	var foundBook bool
	// checking for add or update
	if addNewBook {
		newBook = models.Book{*id, *title, *author, *price}
		books = append(books, newBook)
	} else {
		for i, book := range books {
			if book.Id == *id {
				// replace old values with new ones
				books[i] = models.Book{*id, *title, *author, *price}
				foundBook = true
			}
		}
		// if no book found with mentioned id throws an error
		if !foundBook {
			fmt.Println("Book not found")
			os.Exit(1)
		}
	}
	err := saveBooks(books)
	checkError(err)
	fmt.Println("Book added successfully")
}

func handleDeleteBook(deleteCmd *flag.FlagSet, id *string) {

	deleteCmd.Parse(os.Args[2:])

	if *id == "" {
		fmt.Println("Please provide book --id")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}

	books := getBooks()
	var foundBook bool

	for i, book := range books {
		
		if book.Id == *id {
			// There is no direct delete
			// so creating 2 books structs without the targeted book and appending
			books = append(books[:i], books[i+1:]...)
			foundBook = true
		}
	}
	// if no book found with mentioned id throws an error
	if !foundBook {
		fmt.Println("Book not found")
		os.Exit(1)
	}
	err := saveBooks(books)
	checkError(err)
	fmt.Println("Book deleted successfully")
}
