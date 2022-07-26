package api

import (
	. "hyperlite"
	"net"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/api/router"
	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/logger"
)

type authStruct struct {
	Active bool `json:"active"`
}

var app *App

func RunServer() {
	combinedLogger := logger.GetCombinedLogger()
	combinedLogger.Info("Starting web server")

	routes := router.InitializeRoutes()

	serverDev := viper.GetBool("SERVER.SECURITY")
	port := viper.GetInt("SERVER.PORT")
	fullchainCert := viper.GetString("SERVER.CERT.FULLCHAIN")
	privKeyCert := viper.GetString("SERVER.CERT.PK")
	resolve, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:"+string(port))
	resolveTCP, _ := net.ListenTCP("tcp4", resolve)

	combinedLogger.Info("HTTP about to listen on " + string(port))

	var _ error
	if !serverDev {
		app = NewApp(AppConfig{
			Server: &ServerConfig{
				Host: "0.0.0.0",
				Port: port,
			},
		})
		app.NotFoundContext = func(c *Context) {
			c.JSONError(404, "Endpoint not defined")
		}
		DefineRoutes()
		app.Start()
		//http.Serve(resolveTCP, routes)
	} else {
		http.ServeTLS(resolveTCP, routes, fullchainCert, privKeyCert)
	}

	//combinedLogger.Error("Failed to startup of API server " + err.Error())
}

func DefineRoutes() {
	app.Get("/manager/key", router.GetKeys)
	app.Post("/manager/key", router.CreateKey)
	app.Delete("/manager/key", router.DeleteKey)
}
