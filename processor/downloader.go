package processor

import (
	"log"
	"github.com/chapterzero/gomposer/provider"
)

const tempDirectory = "/tmp"

func downloadPackage(packageName string, provider provider.Provider, version Version, vendorDirectory string) {
	downloadUrl := provider.GetDownloadUrl(packageName, version.Value)
	log.Println(downloadUrl)
}
