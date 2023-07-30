package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"log"
)

var InternalConfig Configuration

type Configuration struct {
	FileName           string `json:"FileName"`
	DbType             string `json:"db_type"`
	DbHost             string `json:"db_host"`
	DbName             string `json:"db_name"`
	DbUser             string `json:"db_user"`
	DbPassword         string `json:"db_password"`
	DbDNS              string `json:"db_dns"`
	SslCertificate     string `json:"ssl_certificate"`
	SslCertificateKey  string `json:"ssl_certificate_key"`
	SslPort            string `json:"ssl_Port"`
	IsSsl              string `json:"is_ssl"`
	Port               string `json:"port"`
	AssetsFileAbsolute string `json:"assets_file_absolute"`
	AssetsFileRelative string `json:"assets_file_relative"`
	LogLevel           string `json:"log_level"`
	LogConsole         string `json:"log_console"`
}

var err error

func InitializeConfiguration() (Configuration, bool) {

	configuration := Configuration{}
	fileName := fmt.Sprintf("config/config.json")
	err = gonfig.GetConf(fileName, &configuration)
	if err != nil {
		fmt.Println("Error in Loading Configuration 1" + err.Error())
		return Configuration{}, false
	}

	log.Println(configuration.FileName)
	err = gonfig.GetConf(configuration.FileName, &configuration)
	if err != nil {
		fmt.Println("Error in Loading Configuration 2" + err.Error())
		return Configuration{}, false
	}
	InternalConfig = configuration
	return configuration, true
}

func GetConfigData(key string) (string, bool) {

	if &InternalConfig == nil {
		log.Println("No Internal Config file")
		return "", false
	}

	switch key {
	case "DbType":
		if InternalConfig.DbType != "" {
			return InternalConfig.DbType, true
		} else {
			return "", false
		}

	case "DbHost":
		if InternalConfig.DbHost != "" {
			return InternalConfig.DbHost, true
		} else {
			return "", false
		}

	case "DbName":
		if InternalConfig.DbName != "" {
			return InternalConfig.DbName, true
		} else {
			return "", false
		}

	case "DbUser":
		if InternalConfig.DbUser != "" {
			return InternalConfig.DbUser, true
		} else {
			return "", false
		}

	case "DbPassword":
		if InternalConfig.DbPassword != "" {
			return InternalConfig.DbPassword, true
		} else {
			return "", false
		}

	case "DbDNS":
		if InternalConfig.DbDNS != "" {
			return InternalConfig.DbDNS, true
		} else {
			return "", false
		}

	default:
		return "", false
	}

}
