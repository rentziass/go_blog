package models

import (
	"errors"
	"log"
)

// A Post is part of our blog
type Post struct {
	id    uint
	title string
	body  string
}

// Getter

// ID ritorna l'id
func (p *Post) ID() uint {
	return p.id
}

// Title ritorna il titolo del post
func (p *Post) Title() string {
	return p.title
}

// Body ritorna il body del post
func (p *Post) Body() string {
	return p.body
}

// Setter

// SetTitle setta titolo
func (p *Post) SetTitle(value string) {
	p.title = value
}

// SetBody setta body del post
func (p *Post) SetBody(value string) {
	p.body = value
}

// Save salva un nuovo Post in DB
func (p Post) Save() {
	db, err := getDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	res, insertErr := db.Exec("INSERT INTO posts(title, body) VALUES ($1, $2)", p.title, p.body)
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	// res, err := stmt.Exec(p.Title(), p.Body())
	// res, err := stmt.Exec()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Print(res)
	defer db.Close()
}

// GetPosts ritorna tutti i post presenti in db
func GetPosts() ([]Post, error) {
	db, err := getDBConnection()

	var posts []Post

	if err == nil {

		rows, err := db.Query(`SELECT id, title, body FROM posts`)

		defer rows.Close()
		for rows.Next() {
			post := Post{}
			scanErr := rows.Scan(&post.id, &post.title, &post.body)
			if scanErr != nil {
				log.Fatal(scanErr)
			} else {
				posts = append(posts, post)
			}
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Printf(rows)
		// defer db.Close()

		return posts, err
	}
	return posts, errors.New("Unable to connect to database")
}

// getter e setter
