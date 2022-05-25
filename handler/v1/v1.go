// SPDX-FileCopyrightText: Copyright (c) 2022 Sidings Media
// SPDX-License-Identifier: MIT

package v1

import (
	"encoding/json"
	"net/http"

	"github.com/SidingsMedia/netbox-prometheus-push-sd/models"
	"github.com/SidingsMedia/netbox-prometheus-push-sd/util"
)

func postTarget(res http.ResponseWriter, req *http.Request) {
	if util.Authenticate(req.Header.Get("X-API-KEY")) {
		res.WriteHeader(http.StatusOK)

		resSuccess := &models.GeneralSuccess{
			Message: "Success",
		}

		data, _ := json.Marshal(resSuccess)
		res.Write(data)

	} else {
		res.WriteHeader(http.StatusUnauthorized)

		resError := &models.GeneralError{
			Code: 401,
			Message: "API key invalid or not present",
		}

		data, _ := json.Marshal(resError)
		res.Write(data)
	}
}

func Target(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	switch req.Method {
		case "POST":
			postTarget(res, req)
		default:
			res.WriteHeader(http.StatusMethodNotAllowed)

			resError := &models.GeneralError{
				Code: 405,
				Message: "Method not allowed",
			}

			data, _ := json.Marshal(resError)
			res.Write(data)
	}
}