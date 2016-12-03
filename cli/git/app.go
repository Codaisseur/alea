package git

import (
	"fmt"
	"os"

	"net/url"
	"strings"

	"github.com/libgit2/git2go"
)

func GetAppFromRemote() string {
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

	url, err := url.Parse(deisRemote.Url())
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.Split(url.EscapedPath(), ".")[0], "/")[1]
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
