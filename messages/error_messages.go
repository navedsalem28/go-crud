package messages

import (
	"encoding/json"
	"go-crud/common"
	"go-crud/logger"
	"io/ioutil"
	"os"
)

var Messages map[int]interface{}

// InitializeMessages from files
func InitializeMessages() {
	Messages = make(map[int]interface{})
	jsonFile, err := os.Open("messages/error_messages.json")
	if err != nil {
		logger.Error(common.InterfaceToString(err.Error()))
		return
	}
	logger.Log("Successfully Opened error.json")

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logger.Error(common.InterfaceToString(err.Error()))
		return
	}
	json.Unmarshal(byteValue, &Messages)
	defer jsonFile.Close()
}

// GetMessageByID func
func GetMessageByID(id int) string {
	return common.InterfaceToString(Messages[id])
}
