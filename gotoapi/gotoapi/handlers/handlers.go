package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"

	"github.com/shiv-ko/Go-tut/gotoapi/models"
)

// GET /hello
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World\n")
}

// POST /article
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// ストリームから取得し、デコード
	// modelsからArticleの型を取得
	var reqArticle models.Article
	// req.BodyをreqArticleにデコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

// GET /article/list
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int
	// クエリパラメータがあるかどうかを確認
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		// intに変換
		page, err = strconv.Atoi(p[0])
		// intに変換できない場合はエラーを返す
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
	//パスパラメータを取得
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	// article1を取得・エンコード・レスポンスに書き込む
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
	// modelsからArticleの型を取得
	var reqArticle models.Article
	// reqArticleをarticleにエンコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

// POST /comment
func CommentArticleHandler(w http.ResponseWriter, req *http.Request) {
	// modelsからArticleの型を取得
	var reqComment models.Comment
	// reqCommentをcommentにエンコード
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}
