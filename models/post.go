package models

import (
	"errors"
	"log"
)

// A Post is part of our blog
type Post struct {
	Id    uint
	Title string
	Body  string
}

// Getter

// // ID ritorna l'id
// func (p *Post) ID() uint {
// 	return p.id
// }
//
// // Title ritorna il titolo del post
// func (p *Post) Title() string {
// 	return p.title
// }
//
// // Body ritorna il body del post
// func (p *Post) Body() string {
// 	return p.body
// }

// Setter

// SetTitle setta titolo
func (p *Post) SetTitle(value string) {
	p.Title = value
}

// SetBody setta body del post
func (p *Post) SetBody(value string) {
	p.Body = value
}

// Save salva un nuovo Post in DB
func (p Post) Save() {
	db, err := getDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	res, insertErr := db.Exec("INSERT INTO posts(title, body) VALUES ($1, $2)", p.Title, p.Body)
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	log.Print(res)
	defer db.Close()
}

// GetPosts ritorna tutti i post presenti in db
func GetPosts() ([]Post, error) {
	db, err := getDBConnection()

	var posts []Post

	if err == nil {
		defer db.Close()
		rows, err := db.Query(`SELECT id, title, body FROM posts`)
		for rows.Next() {
			post := Post{}
			scanErr := rows.Scan(&post.Id, &post.Title, &post.Body)
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

// GetPost returns a post of given id
func GetPost(id int64) (Post, error) {
	db, err := getDBConnection()

	post := Post{}

	if err == nil {
		defer db.Close()
		scanErr := db.QueryRow(`SELECT id, title, body
			FROM posts
			WHERE id=$1`, id).Scan(&post.Id, &post.Title, &post.Body)

		if scanErr != nil {
			log.Fatal(scanErr)
		}
	}

	return post, err
}

// getter e setter
