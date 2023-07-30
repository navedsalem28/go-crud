package main

import (
	"github.com/gorilla/mux"
	"go-crud/config"
	"go-crud/db"
	"go-crud/handler"
	"go-crud/logger"
	"go-crud/messages"
	"net/http"
	"os"
	"time"
)

var ok bool

func main() {
	_, ok = config.InitializeConfiguration()
	if !ok {
		logger.Error("No Config ")
		os.Exit(-1)
	}
	logger.InitLogger(config.InternalConfig.LogLevel, config.InternalConfig.LogConsole)
	logger.Log("Configuration Initialized")
	messages.InitializeMessages()
	if !db.ConnectDB() {
		logger.Error("No DB Connected")
		os.Exit(-1)
	} else {
		logger.Log("DB Connected")
		db.CreateUsersTable()
		defer db.DisconnectDB()
	}
	handler.LoadData()
	routes := UpdateRoute()
	time.Sleep(1 * time.Second)

	logger.Log("******************************************")
	logger.Log("*              Ready to Serve            *")
	logger.Log("******************************************")
	StartServer(routes)

}

func UpdateRoute() *mux.Router {
	Router := mux.NewRouter().StrictSlash(true)
	Router.StrictSlash(true)
	Router.Handle(config.InternalConfig.AssetsFileAbsolute, http.StripPrefix(config.InternalConfig.AssetsFileAbsolute, http.FileServer(http.Dir(config.InternalConfig.AssetsFileRelative))))
	for _, r := range handler.Routes {
		Router.Methods(r.Method).Path(r.Path).Name(r.Name).Handler(r.Handler)
	}
	return Router
}

func StartServer(f *mux.Router) {
	for {
		if config.InternalConfig.IsSsl == "0" {
			e := http.ListenAndServe(":"+config.InternalConfig.Port, f)
			if e != nil {
				logger.Error("HTTP Server not started : " + e.Error())
				time.Sleep(5 * time.Second)
				continue
			}
		} else {
			e := http.ListenAndServeTLS(":"+config.InternalConfig.SslPort, config.InternalConfig.SslCertificate, config.InternalConfig.SslCertificateKey, f)
			if e != nil {
				logger.Error("HTTP Server not started : " + e.Error())
				time.Sleep(5 * time.Second)
				continue
			}
		}
		break
	}
}
