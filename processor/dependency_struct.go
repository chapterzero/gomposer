package processor

import (
	"github.com/chapterzero/gomposer/provider"
)

type Dependency struct {
	Provider        provider.Provider
	FqPackageName   string
	Version         Version
}
