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
	"errors"
	"os/exec"
	"strings"

	"net/url"
	"regexp"
)

var (
	// ErrRemoteNotFound is returned when the remote cannot be found in git
	ErrRemoteNotFound = errors.New("Could not find remote matching app in 'git remote -v'")
	// ErrInvalidRepositoryList is an error returned if git returns unparsible output
	ErrInvalidRepositoryList = errors.New("Invalid output in 'git remote -v'")
	// ErrRemoteNotApp is returned when the remote can't be parsed to an app name
	ErrRemoteNotApp = errors.New("ERROR: Could not resolve app from remote")
)

func GetAppFromRemote() (string, error) {
	uri, err := GetDeisRemoteUri()
	if err != nil {
		return "", err
	}

	return appFromDeisRemote(uri)
}

func GetControllerFromRemote() (string, error) {
	uri, err := GetDeisRemoteUri()
	if err != nil {
		return "", err
	}
	return controllerFromDeisRemote(uri), nil
}

func appFromDeisRemote(remote *url.URL) (string, error) {
	re := regexp.MustCompile("^/([a-zA-Z0-9-_.]+).git")
	matches := re.FindStringSubmatch(remote.EscapedPath())

	if len(matches) == 0 {
		return "", ErrRemoteNotApp
	}

	return string(matches[1]), nil
}

func controllerFromDeisRemote(remote *url.URL) string {
	re := regexp.MustCompile("^([a-z-_]+)(.*)(:\\d+)$")
	return re.ReplaceAllString(remote.Host, "https://services$2")
}

func GetDeisRemoteUri() (*url.URL, error) {
	out, err := exec.Command("git", "remote", "-v").Output()
	emptyUrl, _ := url.Parse("")

	for _, line := range strings.Split(string(out), "\n") {
		// git remote -v contains both push and fetch remotes.
		// They're generally identical, and deis only cares about push.
		if strings.HasSuffix(line, "(push)") {
			parts := strings.Split(line, "\t")
			if len(parts) < 2 {
				return emptyUrl, ErrInvalidRepositoryList
			}

			if parts[0] == "deis" {
				uri, err := url.Parse(strings.Split(parts[1], " ")[0])

				return uri, err
			}
		}
	}

	return emptyUrl, err
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
