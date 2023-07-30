package handler

import (
	"encoding/csv"
	"encoding/json"
	"go-crud/common"
	"go-crud/logger"
	"go-crud/messages"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var categories []Category
var lastCategoryID int
var mutex sync.Mutex

const dataFilePath = "categories.csv"

func LoadData() {
	file, err := os.OpenFile(dataFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logger.Error("Error opening data file:" + err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		logger.Error("Error reading data from file:" + err.Error())
	}

	for _, row := range rows {
		id, _ := strconv.Atoi(row[0])
		categories = append(categories, Category{
			ID:   id,
			Name: row[1],
		})
		lastCategoryID = id
	}
}

func saveData() {
	file, err := os.OpenFile(dataFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Error opening data file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, category := range categories {
		writer.Write([]string{strconv.Itoa(category.ID), category.Name})
	}
	writer.Flush()
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map for response
	ResponseJSON := make(map[string]interface{})
	ResponseJSON["categories"] = categories
	ResponseJSON["message"] = messages.GetMessageByID(0)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
}

func GetSingleCategory(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map for response
	ResponseJSON := make(map[string]interface{})
	ResponseJSON["status"] = 407

	//get category id from URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ResponseJSON["message"] = messages.GetMessageByID(2)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//find category from categories array by id
	category := findCategoryByID(id)
	if category == nil {
		ResponseJSON["message"] = messages.GetMessageByID(8)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	ResponseJSON["category"] = category
	ResponseJSON["message"] = messages.GetMessageByID(0)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map for response
	ResponseJSON := make(map[string]interface{})

	//create RequestJSON map from request body
	RequestJSON, err := common.Input(r)
	if err != nil {
		json.NewEncoder(w).Encode(RequestJSON)
		return
	}

	//get category name from RequestJSON map
	Name, ok := RequestJSON["name"]
	if !ok || common.InterfaceToString(Name) == "" {
		ResponseJSON["message"] = messages.GetMessageByID(11)
		ResponseJSON["status"] = 407
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	//Add new article into articles array
	var category Category
	lastCategoryID++
	category.ID = lastCategoryID
	category.Name = common.InterfaceToString(Name)
	categories = append(categories, category)
	saveData()

	ResponseJSON["message"] = messages.GetMessageByID(12)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

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

	mutex.Lock()
	defer mutex.Unlock()

	//find category index from categories array by id
	index := findCategoryIndexByID(id)
	if index == -1 {
		ResponseJSON["message"] = messages.GetMessageByID(8)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}
	Name, ok := RequestJSON["name"]
	if ok {
		if common.InterfaceToString(Name) == "" {
			ResponseJSON["message"] = messages.GetMessageByID(11)
			ResponseJSON["status"] = 407
			json.NewEncoder(w).Encode(ResponseJSON)
			return
		}
		categories[index].Name = common.InterfaceToString(Name)
		saveData()
	}

	ResponseJSON["message"] = messages.GetMessageByID(4)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
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
	mutex.Lock()
	defer mutex.Unlock()

	//find category from categories array
	index := findCategoryIndexByID(id)
	if index == -1 {
		ResponseJSON["message"] = messages.GetMessageByID(8)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	// Delete category from the categories array
	categories = append(categories[:index], categories[index+1:]...)
	saveData()

	ResponseJSON["message"] = messages.GetMessageByID(13)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
}

func findCategoryByID(id int) *Category {
	for _, category := range categories {
		if category.ID == id {
			return &category
		}
	}
	return nil
}

func findCategoryIndexByID(id int) int {
	for i, category := range categories {
		if category.ID == id {
			return i
		}
	}
	return -1
}
