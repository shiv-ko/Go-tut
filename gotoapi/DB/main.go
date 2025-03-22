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

	// データを挿入する処理

	//記事データ
	article := models.Article{
		Title:    "insert test",
		Content:  "Can I insert data correctly?",
		UserName: "saki",
	}
	//クエリ中に変数を用意
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) 
		values (?, ?, ?, 0, now());
	`
	//変数を埋め込む
	result, err := db.Exec(sqlStr, article.Title, article.Content, article.UserName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 結果の確認
	fmt.Println("記事ID")
	fmt.Println(result.LastInsertId())
	fmt.Println("影響を受けた行数")
	fmt.Println(result.RowsAffected())
	// articleID := 1
	// // SQL文を定義
	// const sqlStr = `
	// 	select *
	// 	from articles
	// 	where article_id = ?;
	// 	`
	// // 実行したいSQL文を引数に指定すると実行結果としてもらえる
	// row := db.QueryRow(sqlStr, articleID)
	// if err := row.Err(); err != nil {
	// 	//データ取得件数が0件だった場合はデータ読み出し処理に移らず終了
	// 	fmt.Println(err)
	// 	return
	// }

	// var article models.Article
	// var createdTime sql.NullTime

	// err = row.Scan(&article.ID, &article.Title, &article.Content,
	// 	&article.UserName, &article.NiceNum, &createdTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if createdTime.Valid {
	// 	article.CreatedAt = createdTime.Time
	// }

	// fmt.Printf("%+v\n", article)
}
