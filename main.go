// SPDX-FileCopyrightText: Copyright (c) 2022 Sidings Media
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"log"
	"net/http"

	v1 "github.com/SidingsMedia/netbox-prometheus-push-sd/handler/v1"
	"github.com/SidingsMedia/netbox-prometheus-push-sd/util"
)

var (
	flagAPIKey string
	flagBindAddress string = ":8081"
	flagPrometheusSDDir string = "/etc/prometheus/sd"
)

// Register API routes
func registerRoutes() {
	http.HandleFunc("/v1/target", v1.Target)
}

func init() {
	// Get flags/env variables
	flag.StringVar(&flagAPIKey, "api-key", util.LookupEnvOrString("NETBOX_PROM_SD_API_KEY", flagAPIKey), "API key to secure endpoints with.")
	flag.StringVar(&flagBindAddress, "bind-address", util.LookupEnvOrString("NETBOX_PROM_SD_BIND_ADDR", flagBindAddress), "Address to bind to. (default ::8081")
	flag.StringVar(&flagPrometheusSDDir, "sd-directory", util.LookupEnvOrString("NETBOX_PROM_SD_DIR", flagPrometheusSDDir), "Directory to store service discovery files. (default /etc/prometheus/sd")
	flag.Parse()

	// Update config
	util.ApiKey = flagAPIKey
	util.BindAddress = flagBindAddress
	util.PrometheusSDDir = flagPrometheusSDDir
}

func main() {
	registerRoutes()
	log.Fatal(http.ListenAndServe(util.BindAddress, nil))
}