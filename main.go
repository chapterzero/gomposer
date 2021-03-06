package main

import (
	"log"
	"path/filepath"
	"strings"
	"os"
	"encoding/json"
	"github.com/chapterzero/gomposer/composer"
	"github.com/chapterzero/gomposer/processor"
)

const ComposerFile = "composer.json"
const VendorDirectory = "govendor"

func main() {
	composerJson, err := readFile()
	if err != nil {
		log.Fatalln("Error when reading composer.json file:", err)
	}

	processor.Process(composerJson, VendorDirectory)
}

func readFile() (composer.ComposerJson, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	var absConfigFile string
	// handle for tmp
	if strings.Contains(dir, "/tmp") {
		absConfigFile = "./" + ComposerFile
	} else {
		absConfigFile = dir + "/" + ComposerFile
	}

	var composerJson composer.ComposerJson

	file, err := os.Open(absConfigFile)
	if err != nil {
		return composerJson, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&composerJson)
	if err != nil {
		return composerJson, err
	}

	return composerJson, nil
}
