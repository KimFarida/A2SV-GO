package controllers

import (
	//"errors"

	"fmt"

	"example.com/library_management/models"
	"example.com/library_management/services"
)

// Now for my subcommands


func CreateMember(name string){
	lib, err:= services.ReadLibraryFile()
	if err != nil{
		fmt.Println("Could not create Memeber ",err)
		return
	}

	id := len(lib.Members) + 1
	newMember := models.Member{
		ID: id,
		Name: name,
		BorrowedBooks: make([]models.Book, 0),

	}

	lib.AddMember(newMember)
	services.WriteToLibraryFile(lib)

}
// Add a new book.
func CreateBook(title string, author string){

	if title == "" || author == ""{
		fmt.Println("Cannot create without Book Title or Author")
		return
	}
	lib, err:= services.ReadLibraryFile()
	
	if err != nil{
		fmt.Println("Could not create Book ",err)
		return
	}

	for _, v := range lib.Books{
		if v.Title == title && v.Author == author{
			fmt.Println("Book already exists")
			return
		}
	}

	newBook := models.Book{
		ID: len(lib.Books) + 1,
		Title: title,
		Author: author,
		Status: models.BookAvailable,
	}
	
	lib.AddBook(newBook)
	services.WriteToLibraryFile(lib)
} 

// Remove an existing book.
func DeleteBook(id int){
	if id < 0{
		fmt.Println("Please Provide ID")
		return
	}
	lib, err:= services.ReadLibraryFile()

	if err != nil{
		fmt.Println("Could not delete Book ",err)
		return
	}
	lib.RemoveBook(id)
	services.WriteToLibraryFile(lib)
}

// Borrow a book.
func BorrowBook(bookID int, memeberID int){
	lib, err:= services.ReadLibraryFile()

	if err != nil{
		println("An error occured:", err)
		return
	}

	err = lib.BorrowBook(bookID, memeberID)

	if err != nil{
		fmt.Println(err)
	}

	services.WriteToLibraryFile(lib)

}

// Return a book.
func ReturnBook(bookID int, memeberID int){
	lib, err:= services.ReadLibraryFile()

	if err != nil{
		println("An error occured:", err)
		return
	}

	err = lib.ReturnBook(bookID, memeberID)

	if err != nil{
		fmt.Println(err)
	}

	services.WriteToLibraryFile(lib)

}
// List all available books.
func ListAvailableBooks()[]models.Book{
	lib, err:= services.ReadLibraryFile()

	if err != nil{
		println("An error occured:", err)
		return nil
	}
	return lib.ListAvailableBooks()
	
}

// List all borrowed books by a member.
func ListBorrowedBooks(memberID int)[]models.Book{
	lib, err:= services.ReadLibraryFile()

	if err != nil{
		println("An error occured:", err)
		return nil
	}
	return lib.ListBorrowed(memberID)
}