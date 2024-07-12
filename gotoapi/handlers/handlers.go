package handlers

import (
  "io"
  "net/http"
)


func HelloHandler(w http.ResponseWriter,req *http.Request){

  if req.Method==http.MethodGet{
    io.WriteString(w,"Hello World\n")
  }
  else {
    http.Error(w,"Invalid Method",http.StatusMethodNotAllowed)
  }
}
func PostArticleHandler(w http.ResponseWriter,req *http.Request){
  io.WriteString(w,"Posting Article...\n")
}
func ArticleListHandler(w http.ResponseWriter,req *http.Request){
  io.WriteString(w,"Article List\n")
}
func BestArticleHandler(w http.ResponseWriter,req *http.Request){
  io.WriteString(w,"Article No.1\n")
}
func NiceArticleHandler(w http.ResponseWriter,req *http.Request){
  io.WriteString(w,"Posting Nicw...\n")
}
func CommentArticleHandler(w http.ResponseWriter,req *http.Request){
  io.WriteString(w,"Posting Comment...\n")
}
