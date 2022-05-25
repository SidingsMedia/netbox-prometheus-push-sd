// SPDX-FileCopyrightText: Copyright (c) 2022 Sidings Media
// SPDX-FileCopyrightText: Copyright (c) 2021 Khanh Ngo <k@ndk.name>
// SPDX-License-Identifier: MIT

package util

import (
	"os"
)

// Lookup specified environment variable. If it is found return its
// value else return the provided default.
// Source:
// https://github.com/ngoduykhanh/wireguard-ui/blob/f43c59c0431bd508df40bdf6b62eaee20287c1de/util/util.go
func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

// Validate the API key
func Authenticate(key string) bool {
	return key == ApiKey
}