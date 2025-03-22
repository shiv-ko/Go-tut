package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shiv-ko/Go-tut/gotoapi/models"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	articleID := 1
	// SQL文を定義
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
		`
	// 実行したいSQL文を引数に指定すると実行結果としてもらえる
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// rowsからデータを読み出して構造体に格納する
	articleArray := make([]models.Article, 0)
	for rows.Next() {
		//　記事データ（１行分）を格納する構造体を定義
		var article models.Article
		// Nullかもしれないデータを格納するための変数を定義
		var createdTime sql.NullTime
		// articleのフィールド内に取得したデータを入れる
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.UserName,
			&article.NiceNum, &createdTime)

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}
		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}
	fmt.Printf("%+v\n", articleArray)
}
