package posts

import (
	"log"

	database "github.com/llucasreis/go-blog/internal/pkg/db/mysql"
	"github.com/llucasreis/go-blog/internal/users"
)

type Post struct {
	ID      string
	Title   string
	Content string
	User    *users.User
}

func (post Post) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Posts(Title, Content) VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(post.Title, post.Content)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Println("Row inserted!")

	return id
}

func GetAll() []Post {
	stmt, err := database.Db.Prepare("select id, title, content from Posts")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)

	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return posts
}
