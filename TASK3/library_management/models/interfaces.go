package models

type LibraryManager interface {
	AddBook(book Book)
	RemoveBook(bookID int)
	BorrowedBook(bookID int, memberID int)error
	ReturnBook(bookID, memberID int)error
	ListAvailableBooks() []Book
	ListBorrowed(memberID int) []Book
	ReserveBook(bookID int, memberID int) error

}