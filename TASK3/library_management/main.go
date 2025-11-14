package main

import (

	"example.com/library_management/concurrency"
	"example.com/library_management/controllers"

	"flag"
	"fmt"
	"os"
	"sync"
)

// Neeed use subCommands and Flags to stimulate console
func main(){
	
	var wg sync.WaitGroup

	wg.Add(1)

	go func(){
		defer wg.Done()
		concurrency.Worker()
	}()
	


	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	rmvCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	reserveCmd := flag.NewFlagSet("reserve",flag.ExitOnError)
	borrowCmd := flag.NewFlagSet("borrow", flag.ExitOnError)
	returnCmd := flag.NewFlagSet("return", flag.ExitOnError)
	avaialbleCmd := flag.NewFlagSet("available-books", flag.ExitOnError)
	borrowedCmd:= flag.NewFlagSet("borrowed-books", flag.ExitOnError)

	var name string
	var title string
	var author string
	var memberID int
	var bookID int

	createCmd.StringVar(&name, "name", "", "Name of Member to create")
	createCmd.StringVar(&title, "title", "", "Title of Book to create")
	createCmd.StringVar(&author, "author", "", "Author of Book to create")

	rmvCmd.IntVar(&bookID, "bookID", 0, "Book ID")

	reserveCmd.IntVar(&bookID, "bookID", 0, "Book ID")
	reserveCmd.IntVar(&memberID, "memberID", 0, "Member ID")

	borrowCmd.IntVar(&bookID, "bookID", 0, "Book ID")
	borrowCmd.IntVar(&memberID, "memberID", 0, "Member ID")

	returnCmd.IntVar(&bookID, "bookID", 0, "Book ID")
	returnCmd.IntVar(&memberID, "memberID", 0, "Member ID")

	borrowedCmd.IntVar(&memberID, "memberID", 0, "Member ID")

	if len(os.Args) < 2{
		fmt.Println("Expected create, borrow or returne subcommands")
	}

	switch os.Args[1]{
		case "create":
			createCmd.Parse(os.Args[2:])
			
			if name != ""{
				controllers.CreateMember(name)
				return
			}
			if title == "" || author == ""{
				fmt.Println("Please Provide a Title and an Author Name if creating a Book OR A name if creating a Member")
				os.Exit(1)
			} 
			controllers.CreateBook(title, author)
			
		case "remove":
			rmvCmd.Parse(os.Args[2:])
			if bookID <= 0{
				fmt.Println("Invalid BookID")
				os.Exit(1)
			}
			controllers.DeleteBook(bookID)

		case "reserve":
			reserveCmd.Parse(os.Args[2:])

			if bookID > 0 && memberID > 0{
				controllers.ReserveBook(bookID, memberID)
			}else{
				for i := 1; i <= 5; i++ {
			
				bookID = 1
				memberID = i + 1 

				controllers.ReserveBook(bookID, memberID)
			}

			}
		case "borrow":
			borrowCmd.Parse(os.Args[2:])

			if bookID <= 0 || memberID <=0{
				fmt.Println("Invalid BookID or MemberID")
				os.Exit(1)
			}
			controllers.BorrowBook(bookID, memberID)
		
		case "return":
			returnCmd.Parse(os.Args[2:])

			if bookID <= 0 || memberID <=0{
				fmt.Println("Invalid BookID or MemberID")
				os.Exit(1)
			}

			controllers.ReturnBook(bookID, memberID)

		case "borrowed-books":
			borrowedCmd.Parse(os.Args[2:])
			if memberID <= 0{
				fmt.Println("Invalid BookID or MemberID")
				os.Exit(1)
			}
			borrowedBooks:= controllers.ListBorrowedBooks(memberID)
			for _, book := range borrowedBooks {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n",book.ID, book.Title, book.Author)
			}
		case "available-books":
			avaialbleCmd.Parse(os.Args[2:])
			books:=controllers.ListAvailableBooks()

			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n",book.ID, book.Title, book.Author)
			}
	}
}
