package main 

import ( 
  "io" 
  "log"
  "net/http" 
  "gotoapi/handlers"
) 
  func main(){
    http.HandleFunc("/hello",handlers.helloHandler)
    http.HandleFunc("/article",handlers.postArticleHandler)
    http.HandleFunc("/article/list",handlers.articleListHandler)
    http.HandleFunc("/article/1",handlers.bestArticleHandler)
    http.HandleFunc("/article/nice",handlers.niceArticleHandler)
    http.HandleFunc("/comment",handlers.commentArticleHandler)
    log.Println("server start at port 8080")
    log.Fatal(http.ListenAndServe(":8080",nil))
        
}
