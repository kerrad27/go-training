package main

import (
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

//go get -u github.com/upper/db/v4/adapter/postgresql

// Set the database credentials using the ConnectionURL type provided by the
// adapter.
var settings = postgresql.ConnectionURL{
	Database: `training`,
	Host:     `localhost:54322`,
	User:     `user_db`,
	Password: `password_db`,
}

//In the event the table contains a column configured to represent automatically-generated values
//like IDs, serials, dates, etc. the omitempty option will have to be added to the db tag:

//type Book struct {
//	ID       uint   `db:"id,omitempty"`
//	Title    string `db:"title"`
//	AuthorID uint   `db:"author_id"`
//}

/*

	CREATE TABLE "books" (
	  "id" INTEGER NOT NULL,
	  "title" VARCHAR NOT NULL,
	  "author_id" INTEGER,
	  CONSTRAINT "books_id_pkey" PRIMARY KEY ("id")
	);

*/
//
//type Author struct {
//	ID        uint   `db:"id,omitempty"`
//	LastName  string `db:"last_name"`
//	FirstName string `db:"first_name"`
//}

/*

	CREATE TABLE "authors" (
	  "id" INTEGER NOT NULL,
	  "first_name" VARCHAR NOT NULL,
	  "last_name" VARCHAR NOT NULL,
	  CONSTRAINT "authors_id_pkey" PRIMARY KEY ("id")
	);

*/

func main() {
	// Use Open to access the database.
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	// The settings variable has a String method that builds and returns a valid
	// DSN. This DSN may be different depending on the database you're connecting
	// to.
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

	// 5/18 Map database records to Go structs

	//booksCol := sess.Collection("books")
	//
	//// Uncomment the following line (and the github.com/upper/db import path) to
	//// write SQL statements to os.Stdout:
	//// db.LC().SetLevel(db.LogLevelDebug)
	//
	//// Find().All() maps all the records from the books collection.
	//var books []Book
	//
	//err = booksCol.Find().All(&books)
	//if err != nil {
	//	log.Fatal("booksCol.Find: ", err)
	//}
	//
	//// Print the queried information.
	//fmt.Printf("Records in the %q collection:\n", booksCol.Title())
	//for i := range books {
	//	fmt.Printf("record #%d: %#v\n", i, books[i])
	//}

	// 6/18 Query database records

	//var books []Book
	//
	//// "books" is a table that already exists in our database.
	//booksTable := sess.Collection("books")
	//
	//// Use Find to create a result set (db.Result).
	//res := booksTable.Find()
	//
	//// The result set can be modified by chaining different db.Result methods
	//// (like Where, And, OrderBy, Select Limit, and Group). Keep in mind that
	//// db.Result is immutable so you'll probably have to reassign the variable
	//// that is holding that object.
	//res = res.OrderBy("-title") // ORDER BY title DESC
	//
	//// Result sets are lazy, meaning that the query will be built or sent to the
	//// database until one of the methods that require database interaction is
	//// used (for example, One or All).
	//
	//if err := res.All(&books); err != nil {
	//	log.Fatal("res.All: ", err)
	//}
	//
	//// The All method copies every single record in the result set into a Go slice.
	//fmt.Printf("Records in the %q table:\n", booksTable.Title())
	//for _, book := range books {
	//	fmt.Printf("%d:\t%q\n", book.ID, book.Title)
	//}
	//fmt.Println("")
	//
	//// Find out how many elements the result set has with Count.
	//total, err := res.Count()
	//if err != nil {
	//	log.Fatal("Count: ", err)
	//}
	//fmt.Printf("There are %d records on %q", total, booksTable.Title())
	//fmt.Println("")
	//
	//// Since result sets are stateless and immutable, they can be reused many
	//// times on different queries.
	//recordsThatBeginWithP := res.And("title LIKE", "T%") // WHERE ... AND title LIKE 'P%'
	//
	//// The original `res` result set is not altered.
	//total1, err := res.Count()
	//if err != nil {
	//	log.Fatal("Count: ", err)
	//}
	//
	//// ... while the new result set is modified.
	//total2, err := recordsThatBeginWithP.Count()
	//if err != nil {
	//	log.Fatal("Count: ", err)
	//}
	//
	//fmt.Printf("There are still %d records on %q\n", total1, booksTable.Title())
	//fmt.Printf("And there are %d records on %q that begin with \"T\"\n", total2, booksTable.Title())

	// 7/18 Query large result sets

	//booksTable := sess.Collection("books")
	//
	//// Order by "id" (descending)
	//res := booksTable.Find().OrderBy("-ID")
	//defer res.Close() // Remember to close the result set.
	//
	//// Next goes over all records one by one. It proves useful when copying large
	//// data sets into a slice is impractical.
	//var book Book
	//for res.Next(&book) {
	//	fmt.Printf("%d:\t%q\n", book.ID, book.Title)
	//}
	//
	//// In the event of a problem Next returns false, that will break the loop and
	//// generate an error (which can be retrieved by calling Err). On the other
	//// hand, when the loop is successfully completed (even if the data set had no
	//// records), Err will be nil.
	//if err := res.Err(); err != nil {
	//	log.Printf("ERROR: %v", err)
	//	log.Fatalf(`SUGGESTION: change OrderBy("-ID") into OrderBy("id") and try again.`)
	//}

	// 8/18 Paginate results

	//authorsCol := sess.Collection("authors")
	//
	//// Create a paginator and sets 10 records by page.
	//res := authorsCol.Find().
	//	OrderBy("last_name", "first_name")
	//
	//p := res.Paginate(3)
	//
	//// Try changing the page number and running the example
	//const pageNumber = 2
	//
	//// Copy all the records from the current page into the customers slice.
	//var authors []Author
	//err = p.Page(pageNumber).All(&authors)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("List of authors (page %d):\n", pageNumber)
	//for i, author := range authors {
	//	fmt.Printf("%d: %q, %q\n", i, author.LastName, author.FirstName)
	//}
	//
	//totalNumberOfEntries, err := p.TotalEntries()
	//if err != nil {
	//	log.Fatal("p.TotalEntries: ", err)
	//}
	//
	//totalNumberOfPages, err := p.TotalPages()
	//if err != nil {
	//	log.Fatal("p.TotalPages: ", err)
	//}
	//
	//fmt.Println("")
	//fmt.Printf("Total entries: %d. Total pages: %d", totalNumberOfEntries, totalNumberOfPages)

	// 10/18 Update, insert, or delete records in a result set

	//booksTable := sess.Collection("books")
	//
	//// This result set includes a single record.
	//res := booksTable.Find(2)
	//
	//// The record is retrieved with the given ID.
	//var book Book
	//err = res.One(&book)
	//if err != nil {
	//	log.Fatal("Find: ", err)
	//}
	//
	//fmt.Printf("Book: %#v", book)
	//
	//// A change is made to a property.
	//book.Title = "New title"
	//
	//fmt.Printf("Book (modified): %#v", book)
	//fmt.Println("")
	//
	//// The result set is updated.
	//if err := res.Update(book); err != nil {
	//	fmt.Printf("Update: %v\n", err)
	//	fmt.Printf("This is OK, we're running on a sandbox with a read-only database.\n")
	//	fmt.Println("")
	//}
	//
	//// The result set is deleted.
	//if err := res.Delete(); err != nil {
	//	fmt.Printf("Delete: %v\n", err)
	//	fmt.Printf("This is OK, we're running on a sandbox with a read-only database.\n")
	//	fmt.Println("")
	//}

	// 11/18 SQL Builder: Select

	//// The Collection / Find / Result syntax was created with compatibility
	//// across all supported databases in mind. However, sometimes it might not be
	//// enough for all your needs, especially when working with complex queries.
	//
	//// In such a case, you can also use the SQL builder.
	//q := sess.SQL().SelectFrom("books")
	//
	//// `q` is a `sqlbuilder.Selector`, you can chain any of its other methods
	//// that also return `Selector`.
	//q = q.OrderBy("title")
	//
	//// Note that queries are immutable, here `p` is a completely independent
	//// query.
	//p := q.Where("title LIKE ?", "T%")
	//
	//// Queries are not compiled nor executed until you use methods like `One` or
	//// `All`.
	//var booksQ, booksP []Book
	//if err := q.All(&booksQ); err != nil {
	//	log.Fatal("q.All: ", err)
	//}
	//
	//// The `Iterator` method is a way to go through large result sets from top to
	//// bottom.
	//booksP = make([]Book, 0, len(booksQ))
	//iter := p.Iterator()
	//var book Book
	//for iter.Next(&book) {
	//	booksP = append(booksP, book)
	//}
	//
	//// Remember to check for error values at the end of the loop.
	//if err := iter.Err(); err != nil {
	//	log.Fatal("iter.Err: ", err)
	//}
	//// ... and to free up any locked resources.
	//if err := iter.Close(); err != nil {
	//	log.Fatal("iter.Close: ", err)
	//}
	//
	//// Listing all books
	//fmt.Printf("All books:\n")
	//for _, book := range booksQ {
	//	fmt.Printf("Book %d:\t%q\n", book.ID, book.Title)
	//}
	//fmt.Println("")
	//
	//// Listing books that begin with P
	//fmt.Printf("Books that begin with \"T\":\n")
	//for _, book := range booksP {
	//	fmt.Printf("Book %d:\t%q\n", book.ID, book.Title)
	//}
	//fmt.Println("")

	// 12/18 SQL Builder: JOIN Queries and Struct Composition

	//The BookAuthorSubject type represents an element that has columns from
	//different tables.
	//type BookAuthorSubject struct {
	//	// The book_id column was added to prevent collisions with the other "id"
	//	// columns from Author and Subject.
	//	BookID uint `db:"book_id"`
	//
	//	Book   `db:",inline"`
	//	Author `db:",inline"`
	//}
	//
	//// This is a query with a JOIN clause that was built using the SQL builder.
	//q := sess.SQL().
	//	Select("b.id AS book_id", "*"). // Note the alias set for book.id.
	//	From("books AS b").
	//	Join("authors AS a").On("b.author_id = a.id").
	//	OrderBy("a.last_name", "b.title")
	//
	//// The JOIN query above returns data from three different tables.
	//var books []BookAuthorSubject
	//if err := q.All(&books); err != nil {
	//	log.Fatal("q.All: ", err)
	//}
	//
	//for _, book := range books {
	//	fmt.Printf("Book %d:\t%s. %q\n", book.BookID, book.Author.LastName, book.Book.Title)
	//}

	// 13/18 SQL Builder: Update, Insert and Delete

	//db.LC().SetLevel(db.LogLevelDebug)
	//
	//var eaPoe Author
	//
	//// We use sqlbuilder.Selector to retrieve the last name "Poe" from the
	//// "authors" table.
	//err = sess.SQL().
	//	SelectFrom("authors").
	//	Where("last_name", "Poe"). // Or Where("last_name = ?", "Poe")
	//	One(&eaPoe)
	//if err != nil {
	//	log.Fatal("Query: ", err)
	//}
	//fmt.Printf("eaPoe: %#v\n", eaPoe)
	//
	//// We use sqlbuilder.Updater to correct the typo in the name "Edgar Allen".
	//res, err := sess.SQL().
	//	Update("authors").
	//	Set("first_name = ?", "Edgar Allan"). // Or Set("first_name", "Edgar Allan").
	//	Where("id = ?", eaPoe.ID).            // Or Where("id", eaPoe.ID)
	//	Exec()
	//if err != nil {
	//	fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
	//}
	//
	////We use sqlbuilder.Inserter to add a new book under "Edgar Allan Poe".
	//book := Book{
	//	Title:    "The Crow",
	//	AuthorID: eaPoe.ID,
	//}
	//res, err = sess.SQL().
	//	InsertInto("books").
	//	Values(book). // Or Columns(c1, c2, c2, ...).Values(v1, v2, v2, ...).
	//	Exec()
	//if err != nil {
	//	fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
	//}
	//if res != nil {
	//	id, _ := res.LastInsertId()
	//	fmt.Printf("New book id: %d\n", id)
	//}

	//// We use sqlbuilder.Deleter to erase the book we've just created (and any
	//// other book with the same name).
	//q := sess.SQL().
	//	DeleteFrom("books").
	//	Where("title", "The Crow")
	//fmt.Printf("Compiled query: %v\n", q)
	//
	//_, err = q.Exec()
	//if err != nil {
	//	fmt.Printf("Query: %v. This is expected on the read-only sandbox\n", err)
	//}

	// 14/18 Raw SQL

	//var eaPoe Author
	//
	//// Query, QueryRow, and Exec are raw SQL methods you can use when db.SQL is
	//// not enough for the complexity of your query.
	//rows, err := sess.SQL().
	//	Query(`SELECT id, first_name, last_name FROM authors WHERE last_name = ?`, "Poe")
	//if err != nil {
	//	log.Fatal("Query: ", err)
	//}
	//
	//// This is a standard query that mimics the API from database/sql.
	//if !rows.Next() {
	//	log.Fatal("Expecting one row")
	//}
	//if err := rows.Scan(&eaPoe.ID, &eaPoe.FirstName, &eaPoe.LastName); err != nil {
	//	log.Fatal("Scan: ", err)
	//}
	//if err := rows.Close(); err != nil {
	//	log.Fatal("Close: ", err)
	//}
	//
	//fmt.Printf("%#v", eaPoe)
	//
	//// Make sure to use Exec or Query, as the case may be.
	//_, err = sess.SQL().
	//	Exec(`UPDATE authors SET first_name = ? WHERE id = ?`, "Edgar Allan", eaPoe.ID)
	//if err != nil {
	//	fmt.Printf("Query: %v. This is expected on the read-only sandbox\n", err)
	//}
	//
	//// The sqlbuilder package provides tools for working with raw sql.Rows, such
	//// as the NewIterator function.
	//rows, err = sess.SQL().
	//	Query(`SELECT * FROM books LIMIT 5`)
	//if err != nil {
	//	log.Fatal("Query: ", err)
	//}
	//
	//// The NewIterator function takes a *sql.Rows value and returns an Iterator.
	//iter := sess.SQL().NewIterator(rows)
	//
	//// This iterator provides methods for going through data, such as All, One,
	//// Next, and the like. If you use Next, remember to use Err and Close too.
	//var books []Book
	//if err := iter.All(&books); err != nil {
	//	log.Fatal("Query: ", err)
	//}
	//
	//fmt.Printf("Books: %#v", books)

	// 17/18 Transactions

	// The `tx` value in the function required by `sess.Tx` is just like `sess`, except
	// it lives within a transaction. This means that if the function returns an
	// error, the transaction will be rolled back.
	//err = sess.Tx(func(tx db.Session) error {
	//	// Anything you set the `tx` variable to execute will be part of the
	//	// transaction.
	//	cols, err := tx.Collections()
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Printf("Cols: %#v\n", cols)
	//
	//	// The booksTable value is valid only within the transaction.
	//	booksTable := tx.Collection("books")
	//	total, err := booksTable.Find().Count()
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Printf("There are %d records in %s\n", total, booksTable.Title())
	//
	//	var books []Book
	//	err = tx.SQL().
	//		SelectFrom("books").Limit(3).OrderBy(db.Raw("RANDOM()")).All(&books)
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Printf("Books: %#v\n", books)
	//
	//	res, err := tx.SQL().
	//		Query("SELECT * FROM books ORDER BY RANDOM() LIMIT 1")
	//	if err != nil {
	//		return err
	//	}
	//
	//	var book Book
	//	err = tx.SQL().NewIterator(res).One(&book)
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Printf("Random book: %#v\n", book)
	//
	//	// If the function returns no error the transaction is committed.
	//	return nil
	//})
	//
	//if err != nil {
	//	fmt.Printf("sess.Tx: ", err)
	//}

}
