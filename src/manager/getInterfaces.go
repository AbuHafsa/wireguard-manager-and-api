package manager

import (
	"errors"
	"fmt"

	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/logger"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func GetInterfaces() ([]*wgtypes.Device, error) { //get interfaces
	client, errInstance := createInstance()
	combinedLogger := logger.GetCombinedLogger()

	if errInstance != nil {
		combinedLogger.Error(fmt.Sprintf("Create instance %s", errInstance))
		return nil, errors.New(errInstance.Error())
	}
	getInterfaces, err := client.Devices()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	closeInstance(client) //release resources and close instance
	return getInterfaces, nil
}
