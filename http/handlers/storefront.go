package handlers

import (
	"ecommerce/database/connection"
	"ecommerce/helpers"
	"ecommerce/lib"
	modelsmysql "ecommerce/models/mysql"
	repositoriesmysql "ecommerce/repositories/mysql"
	"ecommerce/services"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

func CreateStorefront(db *connection.MysqlDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var storefront modelsmysql.StorefrontCreation

		errDecode := json.NewDecoder(r.Body).Decode(&storefront)
		logData := map[string]any{"request": map[string]any{"storefront": storefront}}

		failedMessage := "Storefront creation failed"

		if errDecode != nil {
			logData["response"] = map[string]any{"error": errDecode.Error()}
			helpers.LogError(failedMessage, logData)
			lib.ResponseInvalidRequest(w, nil)
			return
		}

		storefrontService := services.NewStorefrontService(repositoriesmysql.NewStorefrontRepository(db))
		record, errInsert := storefrontService.Create(&storefront)
		if errInsert != nil {
			logData["response"] = map[string]any{"error": errInsert.Error()}
			helpers.LogError(failedMessage, logData)
			lib.ResponseInternalServerError(w, nil)
			return
		}

		successMessage := "Storefront created successfully"
		logData["response"] = record
		helpers.LogInfo(successMessage, logData)
		lib.ResponseCreated(w, record, successMessage)
	}
}

func GetListStorefront(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetListStorefront "+helpers.GetDatetime())
}

func ShowStorefront(db *connection.MysqlDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := GetId(r)

		logData := map[string]any{"request": map[string]any{"id": id}}

		failedMessage := "Storefront not found"

		storefrontService := services.NewStorefrontService(repositoriesmysql.NewStorefrontRepository(db))
		record, err := storefrontService.GetByPrimaryKey(id)
		if err != nil {
			logData["response"] = map[string]any{"error": err.Error()}
			helpers.LogError(failedMessage, logData)

			// update the function key in the log here, will do it later
			if errors.Is(err, gorm.ErrRecordNotFound) {
				lib.ResponseInvalidRequest(w, nil)
			} else {
				lib.ResponseInternalServerError(w, nil)
			}
			return
		}

		successMessage := "Storefront found"
		logData["response"] = record
		helpers.LogInfo(successMessage, logData)
		lib.ResponseOK(w, record, successMessage)
	}
}

func UpdateStorefrontById(db *connection.MysqlDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := GetId(r)
		var storefront modelsmysql.StorefrontUpdate

		errDecode := json.NewDecoder(r.Body).Decode(&storefront)
		logData := map[string]any{"request": map[string]any{"id": id, "storefront": storefront}}

		failedMessage := "Storefront update failed"

		if errDecode != nil {
			logData["response"] = map[string]any{"error": errDecode.Error()}
			helpers.LogError(failedMessage, logData)
			lib.ResponseInvalidRequest(w, nil)
			return
		}

		storefrontService := services.NewStorefrontService(repositoriesmysql.NewStorefrontRepository(db))
		if errUpdate := storefrontService.UpdateById(id, &storefront); errUpdate != nil {
			logData["response"] = map[string]any{"error": errUpdate.Error()}
			helpers.LogError(failedMessage, logData)

			if errors.Is(errUpdate, gorm.ErrRecordNotFound) {
				lib.ResponseInvalidRequest(w, nil)
			} else {
				lib.ResponseInternalServerError(w, nil)
			}
			return
		}

		successMessage := "Storefront updated successfully"
		helpers.LogInfo(successMessage, logData)
		lib.ResponseOK(w, nil, successMessage)
	}
}

func DeleteStorefrontById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteStorefrontById "+helpers.GetDatetime())
}
