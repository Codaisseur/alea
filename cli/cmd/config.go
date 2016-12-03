// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Run:   configure,
	Short: "Configure alea",
}

func init() {
	configCmd.Flags().StringVar(&cfg.controller, "controller", "", "alea controller URL, set it manually if it can't be resolved from the deis git remote, defaults to services.<deis.domain>")
	RootCmd.AddCommand(configCmd)
}

func configure(cmd *cobra.Command, args []string) {
	initConfig()
	if cfg.controller != "" {
		viper.Set("controller", cfg.controller)
		fmt.Println("Controller set to:", cfg.controller)
	}
	writeConfig()
}

func writeConfig() {
	var firstBuffer bytes.Buffer
	if err := toml.NewEncoder(&firstBuffer).Encode(viper.AllSettings()); err != nil {
		// handle error
		panic(err)
	}
	ioutil.WriteFile(cfgFile, []byte(firstBuffer.String()), 0600)
	fmt.Println("Config updated.\n\n", firstBuffer.String())
}
