package main

import (
	"database/sql"
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
	const sqlStr = `
		select title ,contents,username,nice
		from articles;
		`
	// 実行したいSQLぶんを引数に指定すると実行結果としてもらえる
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// rowsからデータを読み出して構造体に格納する
	articleArray := make([]models.Ariticle, 0)
	for rows.Next() {
		var article models.Article
		// articleのフィールドないに取得したデータを入れる
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName,
			&article.NiceNum)
		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}
	fmt.Printf("%+v\n", articleArray)
}
