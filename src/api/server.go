package api

import (
	"net"
	"net/http"

	"github.com/spf13/viper"
	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/api/router"
	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/logger"
)

type authStruct struct {
	Active bool `json:"active"`
}

func RunServer() {
	combinedLogger := logger.GetCombinedLogger()
	combinedLogger.Info("Starting web server")

	routes := router.InitializeRoutes()

	serverDev := viper.GetBool("SERVER.SECURITY")
	port := viper.GetString("SERVER.PORT")
	fullchainCert := viper.GetString("SERVER.CERT.FULLCHAIN")
	privKeyCert := viper.GetString("SERVER.CERT.PK")
	resolve, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:"+port)
	resolveTCP, _ := net.ListenTCP("tcp4", resolve)

	combinedLogger.Info("HTTP about to listen on " + port)

	var err error
	if !serverDev {
		err = http.Serve(resolveTCP, routes)
	} else {
		err = http.ServeTLS(resolveTCP, routes, fullchainCert, privKeyCert)
	}

	combinedLogger.Error("Failed to startup of API server " + err.Error())
}
