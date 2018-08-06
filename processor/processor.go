package processor

import (
	"log"
	"github.com/chapterzero/gomposer/composer"
)

func Process(composerJson composer.ComposerJson, vendorDirectory string) {
	dependencies := dependencyBuilder(composerJson)
	for _, dependency := range dependencies {
		log.Println(dependency.Provider.GetApiUrl(dependency.FqPackageName))
	}
}

// required package may have another requirement
// so we called dependency builder recursively
func dependencyBuilder(composerJson composer.ComposerJson) ([]Dependency) {
	dependencies := make([]Dependency, 0)
	// loop require key
	for fqPackageName, fqVersionName := range composerJson.Require {
		// resolve repository
		provider, err := resolvePackage(fqPackageName, composerJson)
		if err != nil {
			log.Fatalln("Error when processing", fqPackageName, ":" , err)
		}

		version, err := resolveVersion(fqVersionName)
		if err != nil {
			log.Fatalln("Error when processing", fqPackageName, ":", err)
		}

		dComposerJson := provider.GetComposerJson(fqPackageName, version.Value)
		newDependency := Dependency{
			Provider: provider,
			FqPackageName: fqPackageName,
			Version: version,
		}

		dependencies = append(dependencies, newDependency)
		childDependencies := dependencyBuilder(dComposerJson)

		// merge child depenencies
		dependencies = append(dependencies, childDependencies...)
	}

	return dependencies
}
