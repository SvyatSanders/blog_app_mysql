package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
)

// getAllLists — получение всех списков с задачами
func getAllPosts(db *sql.DB) ([]Post, error) {
	rows, err := db.Query("select * from blog_posts.posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]Post, 0, 1)
	for rows.Next() {
		list := Post{}

		err := rows.Scan(&list.ID, &list.Title, &list.Date, &list.Link, &list.Comment)
		if err != nil {
			log.Println(err)
			continue
		}

		res = append(res, list)
	}

	return res, nil
}

// getList — получение поста из ДБ по id
func getPost(db *sql.DB, id string) (Post, error) {
	post := Post{}

	row := db.QueryRow(fmt.Sprintf("select * from blog_posts.posts where posts.id = %v", id))
	err := row.Scan(&post.ID, &post.Title, &post.Date, &post.Link, &post.Comment)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

// updatePost - обновление существующего поста в БД
func updatePost(db *sql.DB, id, title, date, link, comment string) error {
	if len(date) == 0 {
		date = time.Now().Format("2006-01-02 15:04:05")
	}

	if len(title) == 0 {
		return errors.New("EMPTY title")
	}

	res, err := db.Exec("UPDATE blog_posts.posts SET `title` = ?, `date` = ?, `link` = ?, `comment` = ? WHERE `id` = ?",
		title, date, link, comment, id)
	if err != nil {
		return err
	}
	fmt.Printf("result of DB update query: %v", res)

	return nil
}

// deletePost - удаление существующего поста в БД
func deletePost(db *sql.DB, id string) error {
	res, err := db.Exec("DELETE FROM blog_posts.posts WHERE posts.id = ?", id)
	if err != nil {
		return err
	}
	fmt.Printf("result of DB update query: %v", res)

	return nil
}

// createPost - обновление существующего поста в БД
func createPost(db *sql.DB, post Post) error {
	// query := fmt.Sprintf("UPDATE blog_posts.posts SET title = `%s`, date = `%s`, link = `%s`, comment = `%s` WHERE id = '%s'", title, date, link, comment, id)
	if post.Date == "" {
		post.Date = time.Now().Format("2006-01-02 15:04:05")
	}

	if post.Title == "" {
		return errors.New("EMPTY title")
	}

	res, err := db.Exec("INSERT INTO blog_posts.posts (title, date, link, comment) VALUES (?, ?, ?, ?)", post.Title, post.Date, post.Link, post.Comment)
	if err != nil {
		return err
	}
	fmt.Printf("result of DB update query: %v", res)

	return nil
}
