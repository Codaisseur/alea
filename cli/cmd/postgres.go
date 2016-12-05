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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// postgresCmd represents the postgres command
var postgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "Add a PostgreSQL database to your deis app",
	Long: `PostgreSQL database services provide your app with a
separate database and credentials for your app inside the Alea PostgreSQL
service. After the database and credentials are created, your Deis app will
be configured with a DATABASE_URL environment variable that points to the
new database, including the corresponding credentials.`,
	Run: createDatabase,
}

func init() {
	addCmd.AddCommand(postgresCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postgresCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postgresCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func createDatabase(cmd *cobra.Command, args []string) {
	fmt.Print("Creating PostgreSQL database...")

	databaseUrl := requestDatabaseUrl()

	fmt.Printf("Adding new DATABASE_URL to %s...", cfg.app)

	_, err := exec.Command("deis", "config:set", databaseUrl, "-a", cfg.app).Output()
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command("deis", "config").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done!\n\n%s\n", out)
}

func requestDatabaseUrl() string {
	url := fmt.Sprintf("%s/postgres_databases", cfg.controller)

	payload := strings.NewReader(fmt.Sprintf("app=%s", cfg.app))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(" failed!")
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(" done")

	return string(body)
}
