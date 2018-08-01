package processor

import (
	"strings"
	"errors"
)

type Version struct {
	Type     string
	Value    string
}

func resolveVersion(fqVersion string) (Version, error) {
	var emptyVersion Version
	if (strings.HasPrefix(fqVersion, "dev")) {
		split := strings.Split(fqVersion, "-")
		if (len(split) != 2) {
			return emptyVersion, errors.New("Version must have format dev-{branchName}")
		}

		return createBranchVersion(split[1]), nil
	}

	return emptyVersion, errors.New("Could not resolve version: " + fqVersion)
}

func createBranchVersion(branchName string) (Version) {
	version := Version{
		Type: "branch",
		Value: branchName,
	}

	return version
}

