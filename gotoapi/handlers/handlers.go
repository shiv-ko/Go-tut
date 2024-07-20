package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/shiv-ko/Go-tut/gotoapi/models"
)

// GET /hello
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World\n")
}

// POST /article
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "faild to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// GET /article/list
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		// to int
		page, err = strconv.Atoi(p[0])
		// not int
		if err != nil {
			http.Error(w, "Invald query parameter", http.StatusBadRequest)
			return
		}

	} else {
		// クエリパラケータがないときに１に設定する
		page = 1
	}
	// resString := fmt.Sprintf("Article List(page %d)\n", page)
	// io.WriteString(w, resString)

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMSG := fmt.Sprintf("fail to encode json (page%d)\n", page)
		http.Error(w, errMSG, http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// GET /article/{id}
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID%d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// POST /article/nice
func NiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "faild to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// POST /comment
func CommentArticleHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "failed to encode json \n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
