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
	"fmt"
	"os"
	"path/filepath"

	"github.com/codaisseur/alea/git"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

type config struct {
	controller string
	app        string
}

var cfg config

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "alea",
	Short: "Alea command line client",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.alea.yaml)")
	RootCmd.Flags().StringVarP(&cfg.app, "app", "a", "", "deis app name, set it manually if it can't be resolved from the deis git remote")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" { // enable ability to specify config file via flag
		cfgFile = filepath.Join(os.Getenv("HOME"), ".alea.toml")
	}

	if cfg.app == "" {
		cfg.app = git.GetAppFromRemote()
		fmt.Println("App:", cfg.app)
	}

	viper.SetConfigFile(cfgFile)

	viper.SetConfigName(".alea") // name of config file (without extension)
	viper.SetConfigType("toml")  // type of config file (defaults to yaml)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// bind flags
	viper.BindPFlag("controller", RootCmd.PersistentFlags().Lookup("controller"))
	viper.BindPFlag("app", RootCmd.PersistentFlags().Lookup("app"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
