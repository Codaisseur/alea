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
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add service endpoints to your app",
	Long: `Add databases and other services to your app's environment configuration.

Alea currently supports the following services:

  - PostgreSQL databases - alea services add postgresql
	- Mongodb databases - alea services add mongodb
	- Redis endpoints - alea services add redis
	- Memcached endpoints - alea services add memcached

After the service endpoints are created, the endpoints are written to your deis
application config.`,
	Run: addService,
}

func init() {
	servicesCmd.AddCommand(addCmd)
}

func addService(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a service type to add (postgres, mongodb, redis, or memcached)")
		os.Exit(1)
	}
}
