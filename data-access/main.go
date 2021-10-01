package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func nameAlbumsbyArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist: %q %v", name, err)
	}

	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist: %q %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil { //used to check errors in overall query
		return nil, fmt.Errorf("albumsByArtist: %q %v", name, err)
	}
	return albums, nil
}

func getAlbumByID(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsByID %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsByID %d : %v", id, err)
	}
	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title ,artist ,price) VALUES(?,?,?)", alb.Title, alb.Artist, alb.Price)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func main() {
	//Captures connection properties
	dsn := "root:root@tcp(127.0.0.1)/recordings"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	albums, err := nameAlbumsbyArtist("Chainsmokers")
	album, err := getAlbumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	fmt.Printf("Album found : %v\n", album)
	fmt.Printf("ALbums found: %v\n", albums)
	addId, err := addAlbum(Album{
		Title:  "YOu broke me first",
		Artist: "Tata McRae",
		Price:  54.02,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %d", addId)
}
