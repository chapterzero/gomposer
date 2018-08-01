package composer

import (
	"log"
	"encoding/json"
)

func JsonStringToComposerJson(jsonString []byte) (ComposerJson) {
	var composerJson ComposerJson
	err := json.Unmarshal(jsonString, &composerJson)
	if err != nil {
		log.Fatalln("Unable to convert json string to struct:", err)
	}

	return composerJson
}
