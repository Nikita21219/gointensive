package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type ArticleModel struct {
	DB *sql.DB
}

func (a *ArticleModel) Insert(post Article) error {
	query := `INSERT INTO articles (title, text) VALUES($1, $2)`
	row := a.DB.QueryRow(query, post.Title, post.Text)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (a *ArticleModel) GetThreePosts(offset int) ([]Article, error) {
	query := fmt.Sprintf("SELECT id, title, text FROM articles ORDER BY id LIMIT 3 OFFSET %d", offset)
	rows, err := a.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Article
	for rows.Next() {
		var post Article
		err := rows.Scan(&post.Id, &post.Title, &post.Text)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (a *ArticleModel) CountPosts() (int, error) {
	query := `SELECT COUNT(*) from articles`
	row := a.DB.QueryRow(query)
	if err := row.Err(); err != nil {
		return -1, err
	}
	var count int
	err := row.Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (a *ArticleModel) GetArticle(id string) (*Article, error) {
	query := fmt.Sprintf("SELECT title, text FROM articles WHERE id=%s", id)
	row, err := a.DB.Query(query)
	if err != nil {
		return nil, err
	}
	post := new(Article)

	if row.Next() {
		err = row.Scan(&post.Title, &post.Text)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Not found")
	}
	return post, nil
}
