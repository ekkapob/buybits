package query

import (
	"database/sql"
	"github.com/ekkapob/buybits/model"
)

func GetPosts(db *sql.DB, limit, offset int) (posts []model.Post, err error) {
	rows, err := db.Query("SELECT id, owner, title, description, createddate, updateddate, skills, budget FROM posts ORDER BY createddate DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	var post model.Post
	for rows.Next() {
		post = model.Post{}
		rows.Scan(&post.Id, &post.Owner, &post.Title, &post.Description, &post.CreatedDate, &post.UpdatedDate, &post.Skills, &post.Budget)
		posts = append(posts, post)
	}
	return posts, nil
}
