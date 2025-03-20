package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"belajar-sqlc/tutorial"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Belajar menggunakan sql dan goose")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	queries := tutorial.New(conn)
	ctx := context.Background()

	// Tambah Author
	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		log.Fatal("CreateAuthor error:", err)
	}
	fmt.Println("New Author:", insertedAuthor)

	// Ambil Author
	author, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		log.Fatal("GetAuthor error:", err)
	}
	fmt.Println("Get Author:", author)

	// List Author
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		log.Fatal("ListAuthors error:", err)
	}
	fmt.Println("List Authors:", authors)

	// Update Author
	err = queries.UpdateAuthor(ctx, tutorial.UpdateAuthorParams{
		ID:   author.ID,
		Name: "Updated Name",
		Bio:  pgtype.Text{String: "Updated Bio", Valid: true},
	})
	if err != nil {
		log.Fatal("UpdateAuthor error:", err)
	}
	fmt.Println("Author Updated")

	// Hapus Author
	err = queries.DeleteAuthor(ctx, author.ID)
	if err != nil {
		log.Fatal("DeleteAuthor error:", err)
	}
	fmt.Println("Author Deleted")

}
