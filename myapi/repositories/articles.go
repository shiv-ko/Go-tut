// artilesテーブルを操作する関数を実装
package repositories

import (
	"database/sql"
	"github.com/shiv-ko/Go-tut/myapi/models"
)

// POST/article :リクエストボディで受け取った記事を投稿する
// 構造体models.Articleを受け取って、それをDBに挿入する処理を追加
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles(title, content, user_name, nice, created_at) values
		(?, ?, ?,0 , now());
		`
	//構造体models.Articleを受け取って、それをDBに挿入する処理
	result, err := db.Exec(sqlStr, article.Title, article.Content, article.UserName)
	if err != nil {
		return models.Article{}, err
	}
	//挿入した記事のIDを取得
	newArticleID, err := result.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}
	//挿入した記事のIDを構造体に格納
	newArticle := models.Article{
		ID:       int(newArticleID),
		Title:    article.Title,
		Content:  article.Content,
		UserName: article.UserName,
		NiceNum:  0,
	}
	return newArticle, nil
}

// GET/article/list :クエリパラメータで指定されたページ（1ページに五つの記事）の記事一覧を取得する
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, content, user_name, nice from articles
		limit ? offset ?;
		`

	//指定された記事データをDBから取得し、それをmodels.Article型のスライスで返す
	// 記事を複数取得するので、Queryを使う
	//第2引数はlimitで、取得する件数を指定する
	//第3引数はoffsetで、開始位置を指定する
	rows, err := db.Query(sqlStr, 5, (page-1)*5)
	if err != nil {
		return nil, err
	}

	//構造体のスライスを作成
	articleArray := []models.Article{}
	for rows.Next() {
		//構造体を作成
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Content, &article.UserName, &article.NiceNum)
		articleArray = append(articleArray, article)
	}
	return articleArray, nil
}

// GET/article/{id} :指定IDの記事を取得する
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
		`
	//指定IDの記事を取得し、それをmodels.Article型で返す
	// 特定の記事を取得するので、QueryRowを使う
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}
	//構造体を作成
	var article models.Article
	var createdTime sql.NullTime
	//構造体に値を格納
	err := row.Scan(&article.ID, &article.Title, &article.Content, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	//構造体を返す
	return article, nil
}

// POST/article/nice :指定IDの記事にいいねをする
func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
		`

	//指定IDの記事のいいね数を1増やす
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?;`
	//いいね数を1増やす
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}
	//コミット＋コミットにエラーがあった場合はエラーを返す
	if err := tx.Commit(); err != nil {
		return err
	}
	//変数nicenumに現在のいいね数を格納
	return nil
}
