# Library Management System Documentation

## Overview

This document outlines the design and usage of the Library Management System implemented in Go. The system demonstrates fundamental Go features including structs, interfaces, and methods.

## Requirements

### Structs

- **Book**
  - `ID` (int): The unique identifier for the book.
  - `Title` (string): The title of the book.
  - `Author` (string): The author of the book.
  - `Status` (string): The status of the book, which can be "Available" or "Borrowed".

- **Member**
  - `ID` (int): The unique identifier for the member.
  - `Name` (string): The name of the member.
  - `BorrowedBooks` ([]Book): A slice that holds the books borrowed by the member.

### Interfaces

- **LibraryManager**
  - `AddBook(book Book)`: Adds a new book to the library.
  - `RemoveBook(bookID int)`: Removes a book from the library using its ID.
  - `BorrowBook(bookID int, memberID int) error`: Allows a member to borrow a book if it is available.
  - `ReturnBook(bookID int, memberID int) error`: Allows a member to return a borrowed book.
  - `ListAvailableBooks() []Book`: Lists all books currently available in the library.
  - `ListBorrowedBooks(memberID int) []Book`: Lists all books borrowed by a specific member.

## File Structure

```plaintext
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
