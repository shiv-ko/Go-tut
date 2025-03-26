package main

import (
	"database/sql"
	// "dbsample/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	//現在のいいね数を取得するクエリ
	article_id := 1
	const sqlGetNice = `
	select nice from articles where article_id = ?
	`
	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	//変数nicenumに現在のいいね数を格納
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	//いいね数を1増やすクエリ
	const sqlUpdateNice = `
	update articles set nice = ? where article_id = ?
	`

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	if err != nil {
		fmt.Println(err)
		//エラーならロールバックして処理を取り消す
		tx.Rollback()
		return
	}
	//コミットして処理内容を確定
	tx.Commit()
}
