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
	"strings"

	"github.com/spf13/cobra"
)

var memcachedCmd = &cobra.Command{
	Use:   "memcached",
	Short: "Add a Memcached namespace to your deis app",
	Long: `Memcached namespace services provide your app with a
separate namespace for your app inside the Alea Memcached
service. After the namespace is created, your Deis app will
be configured with a MEMCACHED_SERVERS environment variable
that points to the memcached servers, as well as a
MEMCACHED_NAMESPACE that defines the unique namespace for
this app.`,
	Run: createMemcachedDatabase,
}

func init() {
	addCmd.AddCommand(memcachedCmd)
}

func createMemcachedDatabase(cmd *cobra.Command, args []string) {
	fmt.Print("Creating Memcached database...")

	namespaceSettings := strings.Split(RequestServiceUrl("memcached_services"), " ")

	memcachedServer := namespaceSettings[0]
	memcachedNamespace := namespaceSettings[1]

	fmt.Printf("Adding new MEMCACHED_SERVERS and MEMCACHED_NAMESPACE to %s...", cfg.app)

	DeisSilentConfigSet(memcachedServer)
	DeisConfigSet(memcachedNamespace)
}
