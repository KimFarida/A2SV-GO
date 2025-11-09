# 📚 Library Management CLI (Go)

A lightweight command-line application for managing a digital library system.

You can create books or members, list available books, borrow and return books, and remove them --- all from the console.

## ⚙️ Setup

### Build the app

`go build main.go`

### Run the executable

`./main <command> [flags]`


Data is persisted in a library.json file in the project root.
It's automatically created if it doesn't exist.

## 🧾 Commands Overview

Command Description

- create Create a new book or member

- available-books List all books currently available in the library

- borrow Borrow a book for a specific member

- borrowed-books Show all books borrowed by a member

- return Return a borrowed book

- remove Remove a book from the library

🪶 1. Create a Book

Syntax:

`./main create -title="<Title>" -author="<Author>"`

Example:

`./main create -title="Things Fall Apart" -author="Chinua Achebe"`

✅ Result:

A new book is added to the library with a unique ID and marked as available.

👥 2. Create a Member

Syntax:

`./main create -name="<Member Name>"`

Example:

`./main create -name="Bella Hassan"`

✅ Result:

A new member is added to the library's member list.

📖 3. List All Available Books

Syntax:

`./main available-books`

Example Output:

ID: 2, Title: Americana, Author: Chimamanda Adichie

ID: 3, Title: Things Fall Apart, Author: Chinua Achebe

✅ Result:

Displays all books whose status is "available".

📚 4. Borrow a Book

Syntax:

`./main borrow -bookID=<BookID> -memberID=<MemberID>`

Example:

`./main borrow -bookID=2 -memberID=2`

✅ Result:

The book's status changes from "available" to "borrowed".

The book is added to the member's borrowed list.

⚠️ If the member or book doesn't exist, the program prints:

a member with the id does not exist

📘 5. View Borrowed Books

Syntax:

`./main borrowed-books -memberID=<MemberID>`

Example:

`./main borrowed-books -memberID=2`

Output:

`ID: 2, Title: Americana, Author: Chimamanda Adichie`

✅ Result:

Lists all books currently borrowed by that member.

🔁 6. Return a Book

Syntax:

`./main return -bookID=<BookID> -memberID=<MemberID>`

Example:

./main return -bookID=2 -memberID=2

✅ Result:

The book's status changes back to "available".

The book is removed from the member's borrowed list.

❌ 7. Remove a Book

Syntax:

./main remove -bookID=<BookID>

Example:

./main remove -bookID=2

✅ Result:

The specified book is permanently deleted from the library database.



🧩 Example Workflow

# 1. Create books

./main create -title="Americana" -author="Chimamanda Adichie"

./main create -title="Things Fall Apart" -author="Chinua Achebe"

# 2. Create a member

./main create -name="Bella Hassan"

# 3. View all available books

./main available-books

# 4. Borrow a book

./main borrow -bookID=2 -memberID=2

# 5. Check borrowed books for a member

./main borrowed-books -memberID=2

# 6. Return a book

./main return -bookID=2 -memberID=2

# 7. Remove a book

./main remove -bookID=2

🗂️ Data Storage

All operations read/write to library.json.

Example structure:

{

"Books": {

"2": { "ID": 2, "Title": "Americana", "Author": "Chimamanda Adichie", "Status": "available" },

"3": { "ID": 3, "Title": "Things Fall Apart", "Author": "Chinua Achebe", "Status": "available" }

},

"Members": {

"2": { "ID": 2, "Name": "Bella Hassan", "BorrowedBooks": [] }

}

}

💡 Notes

All commands are case-sensitive.

Flags must start with - (e.g., -title, -author, -bookID, -memberID).

library.json must remain in the root folder for persistence.

If the file is missing, the program automatically creates it.