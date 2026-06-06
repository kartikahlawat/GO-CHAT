package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	_ "modernc.org/sqlite"
)

func main() {
	dbPath := "chatapp.db"
	if len(os.Args) > 1 {
		dbPath = os.Args[1]
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database:", dbPath)
	fmt.Println()

	printTables(db)
	printUsers(db)
	printMessages(db)
}

func printTables(db *sql.DB) {
	rows, err := db.Query(`
		SELECT name
		FROM sqlite_master
		WHERE type = 'table'
		ORDER BY name
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Tables")
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println("-", name)
	}
	fmt.Println()
}

func printUsers(db *sql.DB) {
	rows, err := db.Query(`
		SELECT id, username
		FROM users
		ORDER BY id
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users")
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tUsername")
	for rows.Next() {
		var id int
		var username string
		if err := rows.Scan(&id, &username); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(writer, "%d\t%s\n", id, username)
	}
	writer.Flush()
	fmt.Println()
}

func printMessages(db *sql.DB) {
	rows, err := db.Query(`
		SELECT id, sender, receiver, content, created_at
		FROM messages
		ORDER BY id DESC
		LIMIT 50
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Recent Messages")
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tSender\tReceiver\tContent\tCreated At")
	for rows.Next() {
		var id int
		var sender string
		var receiver string
		var content string
		var createdAt string
		if err := rows.Scan(&id, &sender, &receiver, &content, &createdAt); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\n", id, sender, receiver, content, createdAt)
	}
	writer.Flush()
}
