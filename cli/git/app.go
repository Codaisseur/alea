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
	re := regexp.MustCompile("^/([a-zA-Z0-9-_.]+).git")
	matches := re.FindStringSubmatch(uri.EscapedPath())

	if len(matches) == 0 {
		fmt.Println("ERROR: Could not resolve app from remote")
		os.Exit(1)
	}

	return string(matches[1])
}

func GetControllerFromRemote() string {
	uri := GetDeisRemoteUri()
	re := regexp.MustCompile("^([a-z-_]+)(.*)(:\\d+)$")
	return re.ReplaceAllString(uri.Host, "https://services$2")
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
