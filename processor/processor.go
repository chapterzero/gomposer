package processor

import (
	"log"
	"github.com/chapterzero/gomposer/composer"
)

func Process(composerJson composer.ComposerJson) {
	// loop require key
	for fqPackageName, fqVersionName := range composerJson.Require {
		// resolve repository
		provider, err := resolvePackage(fqPackageName, composerJson)
		if err != nil {
			log.Println("Error when processing", fqPackageName, ":" , err)
			continue
		}

		version, err := resolveVersion(fqVersionName)
		if err != nil {
			log.Println("Error when processing", fqPackageName, ":", err)
			continue
		}

		log.Println(provider.GetApiUrl(fqPackageName))
		log.Println(version)
	}
}
