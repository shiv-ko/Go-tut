package main 

import ( 
  "log"
  "net/http" 
  "github.com/shiv-ko/Go-tut/gotoapi/handlers"
  "github.com/gorilla/mux"
) 
  func main(){
    r:=mux.NewRouter()
    r.HandleFunc("/hello",handlers.HelloHandler).Methods(http.MethodGet)
    r.HandleFunc("/article",handlers.PostArticleHandler).Methods(http.MethodPost)
    r.HandleFunc("/article/list",handlers.ArticleListHandler).Methods(http.MethodGet)
    r.HandleFunc("/article/1",handlers.BestArticleHandler).Methods(http.MethodGet)
    r.HandleFunc("/article/nice",handlers.NiceArticleHandler).Methods(http.MethodPost)
    r.HandleFunc("/comment",handlers.CommentArticleHandler).Methods(http.MethodPost)
    log.Println("server start at port 8080")
    log.Fatal(http.ListenAndServe(":8080",r))
}
