package git

import (
	// "reflect"
	"net/url"
	"testing"

	"github.com/arschles/assert"
)

func TestAppFromDeisRemote(t *testing.T) {
	t.Parallel()

	url, err := url.Parse("ssh://git@deis-builder.alea.dev:2222/api.coderunner.git")
	assert.Equal(t, nil, err, "app-from-deis-remote")

	expected := "api.coderunner"
	actual, err := appFromDeisRemote(url)

	assert.Equal(t, expected, actual, "app-from-deis-remote")
	assert.Equal(t, nil, err, "app-from-deis-remote")
}

func TestControllerFromDeisRemote(t *testing.T) {
	t.Parallel()

	url, err := url.Parse("ssh://git@deis-builder.alea.dev:2222/api.coderunner.git")
	assert.Equal(t, nil, err, "controller-from-deis-remote")

	expected := "https://services.alea.dev"
	actual := controllerFromDeisRemote(url)

	assert.Equal(t, expected, actual, "controller-from-deis-remote")
}
