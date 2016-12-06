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

package git

import (
	"fmt"
	"os"

	"net/url"
	"regexp"

	"github.com/libgit2/git2go"
)

func GetAppFromRemote() string {
	uri := GetDeisRemoteUri()
	return appFromDeisRemote(uri)
}

func GetControllerFromRemote() string {
	uri := GetDeisRemoteUri()
	return controllerFromDeisRemote(uri)
}

func appFromDeisRemote(remote *url.URL) string {
	re := regexp.MustCompile("^/([a-zA-Z0-9-_.]+).git")
	matches := re.FindStringSubmatch(remote.EscapedPath())

	if len(matches) == 0 {
		fmt.Println("ERROR: Could not resolve app from remote")
		os.Exit(1)
	}

	return string(matches[1])
}

func controllerFromDeisRemote(remote *url.URL) string {
	re := regexp.MustCompile("^([a-z-_]+)(.*)(:\\d+)$")
	return re.ReplaceAllString(remote.Host, "https://services$2")
}

func GetDeisRemoteUri() *url.URL {
	repo, err := git.OpenRepository(".")
	if err != nil {
		fmt.Println("Not a git repository")
		os.Exit(1)
	}

	remotes, err := repo.Remotes.List()

	if !stringInSlice("deis", remotes) {
		fmt.Println("No deis remote found")
		os.Exit(1)
	}

	deisRemote, err := repo.Remotes.Lookup("deis")
	if err != nil {
		panic(err)
	}

	uri, err := url.Parse(deisRemote.Url())
	if err != nil {
		panic(err)
	}

	return uri
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
