// SPDX-FileCopyrightText: Copyright (c) 2022 Sidings Media
// SPDX-License-Identifier: MIT

package models

type TargetWebhook struct {
	Type string `json:"type"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
	FQDN string `json:"fqdn"`
	Target string `json:"target"`
	Location string `json:"location"`
	Platform string `json:"platform"`
	Status string `json:"status"`
	Role string `json:"role"`
	Vm bool `json:"vm"`
	IPv4 string `json:"ipv4"`
	IPv6 string `json:"ipv6"`
}