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

var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Add a Redis database to your deis app",
	Long: `Redis database services provide your app with a
separate database for your app inside the Alea Redis
service. After the database is created, your Deis app will
be configured with a REDIS_URL environment variable that
points to the new database.`,
	Run: createRedisDatabase,
}

func init() {
	addCmd.AddCommand(redisCmd)
}

func createRedisDatabase(cmd *cobra.Command, args []string) {
	fmt.Print("Creating Redis database...")

	databaseUrl := RequestServiceUrl("redis_services")

	fmt.Printf("Adding new REDIS_URL to %s...", cfg.app)

	DeisConfigSet(databaseUrl)
}
