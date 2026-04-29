package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error Happened ", err)
		os.Exit(1)
	}
}

func getBooks() (books []Book) {
	booksBytes, err := ioutil.ReadFile("./books.json")
	checkError(err)
	err = json.Unmarshal(booksBytes, &books)
	checkError(err)
	return books
}

func saveBooks(books []Book) error {
	booksBytes, err := json.Marshal(books)
	checkError(err)
	err = ioutil.WriteFile("./books.json", booksBytes, 0644)
	return err
}

func handleGetBooks(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])
	if !*all && *id == "" {
		fmt.Println("subcommand --all or --id needed")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		books := getBooks()
		fmt.Printf("Id \t Title \t Author \t Price \t \n")

		for _, book := range books {
			fmt.Printf("%v \t %v \t %v \t %v \t \n", book.Id, book.Title, book.Author, book.Price)
		}
	}

	if *id != "" {
		books := getBooks()
		fmt.Printf("Id \t Title \t Author \t Price \t \n")

		var foundBook bool
		for _, book := range books {
			foundBook = true
			if *id == book.Id {
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", book.Id, book.Title, book.Author, book.Price)
			}
		}
		if !foundBook {
			fmt.Println("Book not found")
			os.Exit(1)
		}
	}
}

func handleAddBook(addCmd *flag.FlagSet, id, title, author, price *string, addNewBook bool) {
	addCmd.Parse(os.Args[2:])

	if *id == "" || *title == "" || *author == "" || *price == "" {
		fmt.Println("Please provide book id, title, author,price")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
	books := getBooks()
	var newBook Book
	var foundBook bool
	if addNewBook {
		newBook = Book{*id, *title, *author, *price}
		books = append(books, newBook)
	} else {
		for i, book := range books {
			if book.Id == *id {
				
				books[i] = Book{*id, *title, *author, *price}
				foundBook = true
			}
		}

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
			books = append(books[:i], books[i+1:]...)
			foundBook = true
		}
	}
	
	if !foundBook {
		fmt.Println("Book not found")
		os.Exit(1)
	}
	err := saveBooks(books)
	checkError(err)
	fmt.Println("Book deleted successfully")
}
