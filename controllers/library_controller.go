package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"github.com/kika1s1/Library-Management-System/models"
	"github.com/kika1s1/Library-Management-System/services"
)

var library = services.NewLibrary()

func StartLibraryConsole() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			addBook(scanner)
		case "2":
			removeBook(scanner)
		case "3":
			borrowBook(scanner)
		case "4":
			returnBook(scanner)
		case "5":
			listAvailableBooks()
		case "6":
			listBorrowedBooks(scanner)
		case "7":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addBook(scanner *bufio.Scanner) {
	var book models.Book

	fmt.Print("Enter Book ID: ")
	scanner.Scan()
	book.ID, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Enter Book Title: ")
	scanner.Scan()
	book.Title = scanner.Text()

	fmt.Print("Enter Book Author: ")
	scanner.Scan()
	book.Author = scanner.Text()

	book.Status = "Available"

	library.AddBook(book)
	fmt.Println("Book added successfully.")
}

func removeBook(scanner *bufio.Scanner) {
	fmt.Print("Enter Book ID to remove: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())

	library.RemoveBook(bookID)
	fmt.Println("Book removed successfully.")
}

func borrowBook(scanner *bufio.Scanner) {
	fmt.Print("Enter Book ID to borrow: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter Member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func returnBook(scanner *bufio.Scanner) {
	fmt.Print("Enter Book ID to return: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter Member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func listAvailableBooks() {
	books := library.ListAvailableBooks()
	fmt.Println("\nAvailable Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func listBorrowedBooks(scanner *bufio.Scanner) {
	fmt.Print("Enter Member ID to list borrowed books: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	books := library.ListBorrowedBooks(memberID)
	fmt.Println("\nBorrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
