package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-crud/common"
	"go-crud/logger"
	"go-crud/messages"
	"net/http"
	"strconv"
)

type Article struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int64  `json:"price"`
}

var Articles []Article
var lastArticleID int

// GetAllArticles to get all articles from  articles array
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map for response
	ResponseJSON := make(map[string]interface{})
	ResponseJSON["articles"] = Articles
	ResponseJSON["message"] = messages.GetMessageByID(0)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
}

// GetSingleArticle to get single article record from articles array
func GetSingleArticle(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map for response
	ResponseJSON := make(map[string]interface{})
	ResponseJSON["status"] = 407

	//get article id from URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ResponseJSON["message"] = messages.GetMessageByID(2)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//find article from articles array
	article := findArticleByID(id)
	if article == nil {
		ResponseJSON["message"] = messages.GetMessageByID(8)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	ResponseJSON["article"] = article
	ResponseJSON["message"] = messages.GetMessageByID(0)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
}

// CreateArticle to create a new article record into articles array
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map for response
	ResponseJSON := make(map[string]interface{})
	ResponseJSON["status"] = 407

	//create RequestJSON map from request body
	RequestJSON, err := common.Input(r)
	if err != nil {
		json.NewEncoder(w).Encode(RequestJSON)
		return
	}

	//get article title from RequestJSON map
	Title, ok := RequestJSON["title"]
	if !ok || common.InterfaceToString(Title) == "" {
		ResponseJSON["message"] = messages.GetMessageByID(9)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//get article price from RequestJSON map
	Price, ok := RequestJSON["price"]
	if !ok || common.InterfaceToInt(Price) == 0 {
		ResponseJSON["message"] = messages.GetMessageByID(10)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//Add new article into articles array
	var article Article
	lastArticleID++
	article.ID = lastArticleID
	article.Title = common.InterfaceToString(Title)
	article.Price = common.InterfaceToInt(Price)
	Articles = append(Articles, article)

	ResponseJSON["message"] = messages.GetMessageByID(12)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)

}

// UpdateArticle to update article record into articles array
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map for response
	ResponseJSON := make(map[string]interface{})
	ResponseJSON["status"] = 407

	//create RequestJSON map from request body
	RequestJSON, err := common.Input(r)
	if err != nil {
		json.NewEncoder(w).Encode(RequestJSON)
		return
	}

	//get article id from URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ResponseJSON["message"] = messages.GetMessageByID(2)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//get article index from articles array by id
	index := findArticleIndexByID(id)
	Title, ok := RequestJSON["title"]
	if ok {
		if common.InterfaceToString(Title) == "" {
			ResponseJSON["message"] = messages.GetMessageByID(9)
			json.NewEncoder(w).Encode(ResponseJSON)
			return
		}
		Articles[index].Title = common.InterfaceToString(Title)
	}

	Price, ok := RequestJSON["price"]
	if ok {
		if common.InterfaceToInt(Price) == 0 {
			ResponseJSON["message"] = messages.GetMessageByID(10)
			json.NewEncoder(w).Encode(ResponseJSON)
			return
		}
		Articles[index].Price = common.InterfaceToInt(Price)
	}

	ResponseJSON["message"] = messages.GetMessageByID(4)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
}

// DeleteArticle to delete article record from articles array
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map for response
	ResponseJSON := make(map[string]interface{})
	ResponseJSON["status"] = 407

	//get article id from URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ResponseJSON["message"] = messages.GetMessageByID(2)
		return
	}

	//get article index from articles array
	index := findArticleIndexByID(id)
	if index == -1 {
		ResponseJSON["message"] = messages.GetMessageByID(8)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}
	// Delete article from the articles array
	Articles = append(Articles[:index], Articles[index+1:]...)

	ResponseJSON["message"] = messages.GetMessageByID(13)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
}

func findArticleByID(id int) *Article {
	for _, article := range Articles {
		if article.ID == id {
			return &article
		}
	}
	return nil
}

func findArticleIndexByID(id int) int {
	for i, article := range Articles {
		if article.ID == id {
			return i
		}
	}
	return -1
}
