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

	"github.com/spf13/cobra"
)

// mongodbCmd represents the mongodb command
var mongodbCmd = &cobra.Command{
	Use:   "mongodb",
	Short: "Add a Mongodb database to your deis app",
	Long: `Mongodb database services provide your app with a
separate database and credentials for your app inside the Alea Mongodb
service. After the database and credentials are created, your Deis app will
be configured with a MONGODB_URL environment variable that points to the
new database, including the corresponding credentials.`,
	Run: createMongoDatabase,
}

func init() {
	addCmd.AddCommand(mongodbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mongodbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mongodbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func createMongoDatabase(cmd *cobra.Command, args []string) {
	fmt.Print("Creating Mongodb database...")

	databaseUrl := RequestServiceUrl("mongodb_services")

	fmt.Printf("Adding new MONGODB_URL to %s...", cfg.app)

	DeisConfigSet(databaseUrl)
}
