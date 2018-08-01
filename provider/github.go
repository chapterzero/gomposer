package provider

import (
	"github.com/chapterzero/gomposer/composer"
	"net/http"
	"encoding/json"
	"encoding/base64"
	"log"
)

const baseUrl = "https://api.github.com/repos";

type Github struct {}
type GithubContentResponse struct {
	Url          string
	DownloadUrl  string
	Content      string
}

func (g Github) GetApiUrl(packageName string) (string) {
	return baseUrl + "/" + packageName
}

func (g Github) GetCommitsUrl(packageName string, sha string) (string) {
	return g.GetApiUrl(packageName) + "/commits/" + sha
}

func (g Github) GetBranchesUrl(packageName string, branchName string) (string) {
	return g.GetApiUrl(packageName) + "/branches/" + branchName
}

func (g Github) GetTagsUrl(packageName string) (string) {
	return g.GetApiUrl(packageName) + "/tags"
}

func (g Github) GetDownloadUrl(packageName string, packageVer string) (string) {
	return g.GetApiUrl(packageName) + "/zipball/" + packageVer
}

func (g Github) GetComposerJson(packageName string, packageVer string) (composer.ComposerJson) {
	url := g.GetApiUrl(packageName) + "/contents/composer.json?ref=" + packageVer
	log.Println("Downloading composer file:", url)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln("Error when getting composer.json from API:", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalln("Error when getting composer.json from API (got non 200 code):", resp.Status)
	}

	defer resp.Body.Close()
	var contentResponse GithubContentResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&contentResponse)

	if err != nil {
		log.Fatalln("Error when reading composer.json file:", err)
	}

	contentDecode, err := base64.StdEncoding.DecodeString(contentResponse.Content)

	if err != nil {
		log.Fatalln("Unable to decode base64 string content:", err)
	}

	return composer.JsonStringToComposerJson(contentDecode)
}
