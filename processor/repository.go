package processor

import (
	"log"
	"strings"
	"errors"
	"net/url"
	"github.com/chapterzero/gomposer/composer"
	"github.com/chapterzero/gomposer/provider"
)

// fqPackageName full qualified package name:
// vendor/packageName
func resolvePackage(fqPackageName string, composerJson composer.ComposerJson) (provider.Provider, error) {
	var emptyProvider provider.Provider

	// validation
	splitString := strings.Split(fqPackageName, "/")
	if (len(splitString) != 2) {
		return emptyProvider, errors.New("Package name need to be vendor/package format")
	}

	// loop repository
	for _, repository := range composerJson.Repositories {
		url, err := url.Parse(repository.Url)
		if err != nil {
			log.Println("Error parsing ", repository.Url, " ", err)
			continue
		}

		pathStriped := url.Path
		if pathStriped[0] == '/' {
			// remove leading slash in front
			pathStriped = pathStriped[1:len(pathStriped)]
		}

		if (pathStriped == fqPackageName) {
			return createProviderFromHost(url.Host)
		}
	}

	return emptyProvider, errors.New("Could not find provider for: " + fqPackageName)
}

func createProviderFromHost(host string) (provider.Provider, error) {
	switch host {
		case "github.com":
			var github provider.Github
			return github, nil
	}

	var emptyProvider provider.Provider
	return emptyProvider, errors.New("Could not find provider for: " + host)
}
