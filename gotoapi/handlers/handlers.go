package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World\n")
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.query()
	var page int

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		//to int
		page, err = strconv.Atoi(p[0])
		// not int
		if err != nil {
			http.Error(w, "Invald query parameter", http.StatusBadRequerst)
			return
		}

	} else {
		//クエリパラケータがないときに１に設定する
		page = 1
	}
	io.WriteString(w, "Article List\n")
}
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	//varsはパスパラメタをMap形式で返す
	//Atoi=string→int
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	///数字じゃない場合のエラー
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}
func NiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}
func CommentArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
