package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func connect() (*sql.DB, bool) {

	cfg := mysql.Config{
		User:   os.Args[1],
		Passwd: os.Args[2],
		Net:    os.Args[3],
		Addr:   os.Args[4],
		DBName: os.Args[5],
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
	return db, true
}
func main() {

	database, _ := connect()

	printAlbums(albumByArtist("John Coltrane", database))
	printAlbum(albumByID(1, database))

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}, database)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

}

func printAlbums(albums []Album, err error) {
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)
}

func printAlbum(album Album, err error) {
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album found: %v", album)
}

func albumByArtist(name string, database *sql.DB) ([]Album, error) {
	var albums []Album

	rows, err := database.Query("SELECT * FROM album WHERE artist = ?;", name)

	if err != nil {
		return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}

		albums = append(albums, alb)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func albumByID(id int64, database *sql.DB) (Album, error) {
	var alb Album
	row := database.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("Album by ID %d: no such album id", id)
		}

		return alb, fmt.Errorf("Album by ID %d: %v", id, err)
	}

	return alb, nil
}

func addAlbum(alb Album, database *sql.DB) (int64, error) {
	result, err := database.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
