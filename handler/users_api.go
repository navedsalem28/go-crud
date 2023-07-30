package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-crud/common"
	"go-crud/db"
	"go-crud/logger"
	"go-crud/messages"
	"net/http"
	"strconv"
)

// GetAllUser to get all users records from database
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map for response
	ResponseJSON := make(map[string]interface{})

	//find users from mySQL database
	params := make([]interface{}, 0)
	AllUsers, _ := db.GetAllRows("SELECT * FROM `users` ", params)
	ResponseJSON["users"] = AllUsers

	ResponseJSON["message"] = messages.GetMessageByID(0)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
	return

}

// GetSingleUser to get single user record from database
func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map for response
	ResponseJSON := make(map[string]interface{})
	ResponseJSON["status"] = 407

	//get user id from URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ResponseJSON["message"] = messages.GetMessageByID(2)
		return
	}

	//find user from mySQL database
	params := make([]interface{}, 1)
	params[0] = id
	User, ok := db.GetSingleRow("SELECT * FROM `users` where `id`=? ", params)
	if !ok {
		ResponseJSON["message"] = messages.GetMessageByID(8)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	ResponseJSON["user"] = User
	ResponseJSON["message"] = messages.GetMessageByID(0)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
}

// CreateUser to create a new user record into database
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	//get user name from RequestJSON map
	Name, ok := RequestJSON["name"]
	if !ok || common.InterfaceToString(Name) == "" {
		ResponseJSON["message"] = messages.GetMessageByID(5)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//get user email from RequestJSON map
	Email, ok := RequestJSON["email"]
	if !ok || !common.IsEmailValid(common.InterfaceToString(Email)) {
		ResponseJSON["message"] = messages.GetMessageByID(14)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//find user from mySQL database
	params := make([]interface{}, 1)
	params[0] = Email
	_, ok = db.GetSingleRow("SELECT * FROM `users` where `email`=? ", params)
	if ok {
		ResponseJSON["message"] = messages.GetMessageByID(15)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//get user phone from RequestJSON map
	Phone, ok := RequestJSON["phone"]
	if !ok || common.InterfaceToString(Phone) == "" {
		ResponseJSON["message"] = messages.GetMessageByID(7)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//insert user record onto database
	params = make([]interface{}, 3)
	params[0] = Name
	params[1] = Email
	params[2] = Phone
	_, ok = db.UpdateDB("Insert into users (`name`,`email`,`phone`) values (?,?,?) ", params)
	if !ok {
		ResponseJSON["message"] = messages.GetMessageByID(1)
		ResponseJSON["status"] = 500
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	ResponseJSON["message"] = messages.GetMessageByID(12)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
	return
}

// UpdateUser to update a user record into database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	//get user id from URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ResponseJSON["message"] = messages.GetMessageByID(2)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//find user from mySQL database
	params := make([]interface{}, 1)
	params[0] = id
	_, ok := db.GetSingleRow("SELECT * FROM `users` where `id`=? ", params)
	if !ok {
		ResponseJSON["message"] = messages.GetMessageByID(8)
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//get user params to update
	params = make([]interface{}, 0)
	updateData := ""
	Name, ok := RequestJSON["name"]
	if ok {
		if common.InterfaceToString(Name) == "" {
			ResponseJSON["message"] = messages.GetMessageByID(5)
			json.NewEncoder(w).Encode(ResponseJSON)
			return
		}
		updateData = updateData + " `name`=?"
		params = append(params, Name)
	}

	Email, ok := RequestJSON["email"]
	if ok {
		if !common.IsEmailValid(common.InterfaceToString(Email)) {
			ResponseJSON["message"] = messages.GetMessageByID(14)
			json.NewEncoder(w).Encode(ResponseJSON)
			return
		}
		if updateData == "" {
			updateData = updateData + "`email`=?"
			params = append(params, Name)
		} else {
			updateData = updateData + ",`email`=?"
			params = append(params, Email)
		}
	}

	Phone, ok := RequestJSON["phone"]
	if ok {
		if common.InterfaceToString(Phone) == "" {
			ResponseJSON["message"] = messages.GetMessageByID(7)
			json.NewEncoder(w).Encode(ResponseJSON)
			return
		}
		if updateData == "" {
			updateData = updateData + "`phone`=?"
			params = append(params, Phone)
		} else {
			updateData = updateData + ",`phone`=?"
			params = append(params, Phone)
		}
	}
	params = append(params, id)

	//update user record into database
	if updateData != "" {
		_, ok = db.UpdateDB("Update  users Set "+updateData+" where `id`= ?", params)
		if !ok {
			ResponseJSON["message"] = messages.GetMessageByID(1)
			ResponseJSON["status"] = 500
			json.NewEncoder(w).Encode(ResponseJSON)
			return
		}
	}

	ResponseJSON["message"] = messages.GetMessageByID(4)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
	return

}

// DeleteUser to delete a user record from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	logger.Debug(r.Proto + " Request From: " + r.RemoteAddr + "\tURL: " + r.URL.Path)

	//add header
	common.AddRestHeader(w)

	//create ResponseJSON map from request body
	ResponseJSON := make(map[string]interface{})

	//get user id from URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ResponseJSON["message"] = messages.GetMessageByID(2)
		ResponseJSON["status"] = 407
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	//delete user from database
	params := make([]interface{}, 1)
	params[0] = id
	_, ok := db.UpdateDB("DELETE FROM users WHERE id = ? ", params)
	if !ok {
		ResponseJSON["message"] = messages.GetMessageByID(1)
		ResponseJSON["status"] = 500
		json.NewEncoder(w).Encode(ResponseJSON)
		return
	}

	ResponseJSON["message"] = messages.GetMessageByID(13)
	ResponseJSON["status"] = 200
	json.NewEncoder(w).Encode(ResponseJSON)
	return

}
