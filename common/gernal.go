package common

import (
	"encoding/json"
	"go-crud/logger"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// InterfaceToString (interface{}) get value of an interface return as string {
func InterfaceToString(ThisInterface interface{}) string {
	if ThisInterface == nil {
		log.Println("InterfaceToString :: NIL value passed for conversion")
		return ""
	}
	switch ThisInterface.(type) {
	case int:
		return strconv.Itoa(ThisInterface.(int))
	case int64:
		return strconv.FormatInt(ThisInterface.(int64), 10)
	case float64:
		TmpStr := strconv.FormatFloat(ThisInterface.(float64), 'f', 10, 64)
		if TmpStr[(len(TmpStr)-10):] == "0000000000" {
			return TmpStr[:(len(TmpStr) - 11)]
		}
		return TmpStr
	default:
		return ThisInterface.(string)
	}
}

// InterfaceToInt (interface{}) get value of interface and return as int64
func InterfaceToInt(ThisInterface interface{}) int64 {
	switch ThisInterface.(type) {
	case int:
		return ThisInterface.(int64)
	case int64:
		return ThisInterface.(int64)
	case float64:
		//return ThisInterface.(int64)
		return int64(ThisInterface.(float64))
	default:
		val, _ := strconv.ParseInt(ThisInterface.(string), 10, 64)
		return val
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/404.html")
}

func User(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/user.html")
}
func Article(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/article.html")
}

func Category(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/category.html")
}

func Input(r *http.Request) (map[string]interface{}, error) {
	ResponseJSON := make(map[string]interface{})
	BodyText, err := ioutil.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)
	if err != nil {
		ResponseJSON["message"] = "Invalid Request Body"
		ResponseJSON["status"] = 406
		logger.Error("controlCommand: can't ready request body text, Error: " + err.Error())
		return ResponseJSON, err
	}
	logger.Log("Request Body Text:\t" + string(BodyText))
	err = json.Unmarshal(BodyText, &ResponseJSON)
	if err != nil {
		ResponseJSON["message"] = "Invalid JSON"
		ResponseJSON["status"] = 406
		logger.Error("controlCommand: Failed Parsing JSON: " + err.Error())
		return ResponseJSON, err
	}

	return ResponseJSON, err
}

func AddRestHeader(w http.ResponseWriter) {
	w.Header().Set("Server", "Mux Web Server")
	w.Header().Set("DevAdmin", "info@webserver.com")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Robots-Tag", "noindex")
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// IsEmailValid checks if the email provided passes the required structure and length.
func IsEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
