package router

import (
	"encoding/json"
	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/db"
	"hyperlite"
	"net/http"
)

type BadRequestResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetKeys(c *hyperlite.Context) {
	_, keys := db.ReturnKeys()
	c.JSON(http.StatusOK, keys)
}

type CreateKeyDto struct {
	PresharedKey string
	PublicKey    string
	BWLimit      int64
	SubExpiry    string
	IPIndex      int
}

func CreateKey(c *hyperlite.Context) {
	body, _ := json.Marshal(c.Body)
	badRequest := false
	//combinedLogger := logger.GetCombinedLogger()

	preKey := CreateKeyDto{}
	json.Unmarshal(body, &preKey)
	if preKey.PresharedKey == "" || preKey.PublicKey == "" {
		c.JSON(http.StatusBadRequest, BadRequestResponse{
			Code:    11,
			Message: "presharedKey and publicKey must be filled",
		})
		badRequest = true
	} else if preKey.BWLimit < 0 {
		c.JSON(http.StatusBadRequest, BadRequestResponse{
			Code:    12,
			Message: "bandwidth cannot be negative",
		})
		badRequest = true
	} else if preKey.SubExpiry == "" {
		c.JSON(http.StatusBadRequest, BadRequestResponse{
			Code:    13,
			Message: "subscription expiry must be filled",
		})
		badRequest = true
	} else if preKey.IPIndex < 0 {
		c.JSON(http.StatusBadRequest, BadRequestResponse{
			Code:    14,
			Message: "IP index must be greater than one",
		})
		badRequest = true
	}

	if !badRequest {
		_, response := db.CreateKey(
			preKey.PublicKey,
			preKey.PresharedKey,
			preKey.BWLimit,
			preKey.SubExpiry,
			preKey.IPIndex)
		c.JSON(http.StatusOK, response)
	}
}

type DeleteKeyDto struct {
	KeyID string
}

func DeleteKey(c *hyperlite.Context) {
	body, _ := json.Marshal(c.Body)
	badRequest := false
	//combinedLogger := logger.GetCombinedLogger()

	data := DeleteKeyDto{}
	json.Unmarshal(body, &data)

	if data.KeyID == "" {
		c.JSON(http.StatusBadRequest, BadRequestResponse{
			Code:    31,
			Message: "keyID needs to be filled",
		})
		badRequest = true
	}

	if !badRequest {
		_, response := db.DeleteKey(data.KeyID)
		c.JSON(http.StatusOK, response)
	}
}
