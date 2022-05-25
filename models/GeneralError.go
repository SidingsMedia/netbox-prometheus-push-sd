// SPDX-FileCopyrightText: Copyright (c) 2022 Sidings Media
// SPDX-License-Identifier: MIT

package models

type GeneralError struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
}