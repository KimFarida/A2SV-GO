package concurrency
import (
	"fmt"
	"example.com/library_management/controllers"
	"example.com/library_management/services"
	"example.com/library_management/models"
	"time"
)

func getLibrary() *models.Library {
    lib, err := services.ReadLibraryFile()
	if err != nil{
		fmt.Println("error loading library data: ", err)
		return  nil
	}
	return lib
}


func autoCancelReservation(bookID int){
	time.Sleep(5 * time.Second)

	lib := getLibrary()

	lib.Mu.Lock()

	if _, ok:= lib.ReservedBooks[bookID]; ok{
		delete(lib.ReservedBooks, bookID)
		fmt.Println("Reservation for book", bookID, "has been cancelled")
	}

	lib.Mu.Unlock()
}

func Worker(){
	fmt.Println("Worker started, listening for reservations...")
	println("==================================================")

	for req := range controllers.ReservationChannel{

		lib := getLibrary()
		lib.Mu.Lock()

		if _, ok := lib.ReservedBooks[req.BookID]; ok{
			fmt.Println("Book ",req.BookID, " This book has already been reserved")
		}else{

			lib.ReservedBooks[req.BookID] = req.MemberID
			fmt.Println("Book ",req.BookID, " This book has been reserved by member", req.MemberID)

			// start autocancellation process
			go autoCancelReservation(req.BookID)
			
		}
		lib.Mu.Unlock()
	}
}
