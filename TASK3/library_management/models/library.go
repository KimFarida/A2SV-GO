package models

import (
	"errors"
	"slices"
	"sync"
)


type Library struct{
	Books map[int]Book
	Members map[int]Member
	ReservedBooks map[int]int 
	Mu sync.Mutex

}

func NewLibrary()*Library{
	return &Library{
		Books: make(map[int]Book),
		Members: make(map[int]Member),
		ReservedBooks : make(map[int]int),
	}
}

// For Testing
func (lib *Library) AddMember(member Member){
	id := member.ID
	allMemebers := lib.Members
	if _, ok := allMemebers[id]; !ok {
		allMemebers[id] = member
	}
	lib.Members = allMemebers

}

// AddBook: Adds a new book to the library.
// I can add a new book to my library, I should check if it exists
func (lib *Library) AddBook(book Book){
	id := book.ID
	allBooks := lib.Books
	if _, ok := allBooks[id]; !ok {
		
		allBooks[id] = book
	}
	lib.Books = allBooks

}
// RemoveBook: Removes a book from the library by its ID.
func (lib *Library) RemoveBook(bookID int){

	_, ok  := lib.Books[bookID]
	if ok {
		delete(lib.Books, bookID)
	}


}

func (lib *Library)BorrowBook(bookID, memberID int)error{
	member, ok := lib.Members[memberID]

	if !ok {
		return errors.New("a memeber with the id does not exist")
	}

	book, ok := lib.Books[bookID]

	if !ok{
		return  errors.New("a book with the id does not exist")
	}

	if book.Status == BookBorrowed{
		return errors.New("this book has already been borrowed")
	}

	// I need to update the book status and remove that book id from the members list
	book.Status = BookBorrowed
	lib.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.Members[memberID] = member

	return nil

}

// BorrowBook: Allows a member to borrow a book if it is available.
// check if the member exists or the book exists
func (lib *Library) ReturnBook(bookID int, memberID int)error{;
	
	book, ok := lib.Books[bookID]

	if !ok{
		return  errors.New("a book with the id does not exist")
	}

	// I need to update the book status and remove that book id from the members list
	
	lib.Mu.Lock()
	defer lib.Mu.Unlock()

	member, ok := lib.Members[memberID]
	if !ok {
		return errors.New("a memeber with the id does not exist")
	}

	member.BorrowedBooks = slices.DeleteFunc(member.BorrowedBooks , func(book Book) bool {
		return book.ID == bookID
	})

	lib.Members[bookID] = member

	book.Status = BookAvailable
	lib.Books[bookID] = book

	

	return nil

}


// ListAvailableBooks() []Book
func (lib *Library)ListAvailableBooks() []Book{
	availableBooks := make([]Book, 0)
	for _, v := range lib.Books{
		if v.Status == BookAvailable{
			availableBooks = append(availableBooks, v)
		}
	}
	return availableBooks
}


// ListBorrowed(memberID int) []Book
func (lib *Library) ListBorrowed(memberID int) []Book{
	member, ok := lib.Members[memberID]

	if ok{
		return member.BorrowedBooks
	}

	return nil
}

// If the book is available, reserve it and process borrowing asynchronously.
// If already reserved, return an error.
// Once I reserve a book, I should start a time based go routine I guess
// Once it returns to the channel, then i unreserve the book
