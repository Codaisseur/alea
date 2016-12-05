// Copyright © 2016 Codaisseur BV <info@codaisseur.com>
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
	RootCmd.AddCommand(configCmd)
}

func configure(cmd *cobra.Command, args []string) {
	if cfg.controller != "" {
		fmt.Println("Controller set to:", viper.GetString("controller"))
		writeConfig()
	}
	fmt.Println("\nConfig:\n----------------")
	fmt.Println(getConfig())
}

func writeConfig() {
	ioutil.WriteFile(cfgFile, []byte(tomlConfig()), 0600)
	fmt.Println("Config updated.\n")
}

func tomlConfig() string {
	var buffer bytes.Buffer
	if err := toml.NewEncoder(&buffer).Encode(viper.AllSettings()); err != nil {
		// handle error
		panic(err)
	}

	return buffer.String()
}

func getConfig() string {
	dat, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		panic(err)
	}

	return string(dat)
}
