package main 

import ( 
  "log"
  "net/http" 
  "github.com/shiv-ko/Go-tut/gotoapi/handlers"
) 
  func main(){
    http.HandleFunc("/hello",handlers.HelloHandler)
    http.HandleFunc("/article",handlers.PostArticleHandler)
    http.HandleFunc("/article/list",handlers.ArticleListHandler)
    http.HandleFunc("/article/1",handlers.BestArticleHandler)
    http.HandleFunc("/article/nice",handlers.NiceArticleHandler)
    http.HandleFunc("/comment",handlers.CommentArticleHandler)
    log.Println("server start at port 8080")
    log.Fatal(http.ListenAndServe(":8080",nil))
}
