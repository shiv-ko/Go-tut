package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shiv-ko/go-tut/gotoapi/models"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World\n")
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsondData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "faild to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)

}
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		//to int
		page, err = strconv.Atoi(p[0])
		// not int
		if err != nil {
			http.Error(w, "Invald query parameter", http.StatusBadRequest)
			return
		}

	} else {
		//クエリパラケータがないときに１に設定する
		page = 1
	}
	resString := fmt.Sprintf("Article List(page %d)\n", page)
	io.WriteString(w, resString)
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
