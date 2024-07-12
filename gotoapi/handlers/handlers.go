package handlers

import (
  "io"
  "net/http"
  "fmt"
)


func HelloHandler(w http.ResponseWriter,req *http.Request){
  io.WriteString(w,"Hello World\n")
}
func PostArticleHandler(w http.ResponseWriter,req *http.Request){
  io.WriteString(w,"Posting Article...\n")
}
func ArticleListHandler(w http.ResponseWriter,req *http.Request){
  io.WriteString(w,"Article List\n")
}
func ArticleDetailHandler(w http.ResponseWriter,req *http.Request){
  articleID,err:=strconv.Atoi(mux.Vars(req)["id"])
  if err!=nil{
    http.Error(w,"Invalid query parameter",http.StatusBadRequest)
    return
  }
  resString:=fmt.Sprintf("Article No.%d\n",articleID)
  io.WriteString(w,resString)
}
func NiceArticleHandler(w http.ResponseWriter,req *http.Request){
  io.WriteString(w,"Posting Nice...\n")
}
func CommentArticleHandler(w http.ResponseWriter,req *http.Request){
  io.WriteString(w,"Posting Comment...\n")
}
