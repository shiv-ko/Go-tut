// commentsテーブルを操作する関数を実装
package repositories

import (
	"database/sql"

	"github.com/shiv-ko/Go-tut/myapi/models"
)

// POST/comment :リクエストボディで受け取ったコメントを投稿する
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments(article_id, message, created_at) values
		(?, ?, now());
		`
	//構造体models.Commentを受け取って、それをDBに挿入する処理
	var newComment models.Comment
	newComment.ArticleID = comment.ArticleID
	newComment.Message = comment.Message
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}
	//挿入したコメントのIDを取得
	newCommentID, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}
	newComment.CommentID = int(newCommentID)
	return newComment, nil
}

// GET/comment/{id} :指定IDのコメント一覧を取得する
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
		`
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := []models.Comment{}
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		commentArray = append(commentArray, comment)
	}
	//指定された記事IDのコメントデータをDBから取得し、それをmodels.Comment型のスライスで返す
	return commentArray, nil
}
