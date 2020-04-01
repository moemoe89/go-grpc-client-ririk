//
//  Practicing gRPC Client
//
//  Copyright © 2020. All rights reserved.
//

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

// ConfigurationModel represent the configuration model
type ConfigurationModel struct {
	RunMode    string `json:"run_mode"`
	Port       string `json:"port"`
	GrpcServer string `json:"grpc_server"`
}

var (
	// Configuration represent the variable of configuration model
	Configuration = &ConfigurationModel{}
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	basepath = strings.Replace(basepath, "config", "", -1)
	file := basepath + "config.json"
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("Failed to load auth configuration file: %s", err.Error()))
	}

	err = json.Unmarshal(raw, Configuration)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse auth configuration file: %s", err.Error()))
	}
}
