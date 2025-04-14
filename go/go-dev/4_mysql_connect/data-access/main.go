package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-faker/faker/v4"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float64
}

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	defer fmt.Println("Closing database connection.")
	defer db.Close()

	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
	fmt.Printf("Album found: %v\n", alb)

	newAlbum := Album{
		Title:  faker.Name(),
		Artist: faker.Name(),
		Price:  10.0,
	}

	_, err_result := insertRow(newAlbum)
	fmt.Println(err_result)

}

func insertRow(album Album) (sql.Result, error) {
	var result sql.Result
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?); SELECT LAST_INSERT_ID();", album.Title, album.Artist, album.Price)
	fmt.Println("Insert result: %v\n", result)
	fmt.Println("Insert result: %v\n", err)
	return result, err
}

func albumByID(id int64) (Album, error) {
	var album Album
	result := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := result.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumById %d: no such album", id)
		}
		return album, fmt.Errorf("albumById %d: %v", id, err)
	}

	return album, nil

}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
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
