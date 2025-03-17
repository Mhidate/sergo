package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Belajar menggunakan sql dan goose")
	connStr := os.Getenv("DATABASE_URL")
	fmt.Println(connStr)
	if connStr == "" {
		fmt.Errorf("DATABASE_URL tidak ditemukan dalam environment variable")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Errorf("gagal membuka koneksi database: %v", err)
	}

	// Coba koneksi ke database
	if err := db.Ping(); err != nil {
		fmt.Errorf("gagal koneksi ke database: %v", err)
	} else {
		fmt.Println("koneksi ke database berhasil")
	}

}
