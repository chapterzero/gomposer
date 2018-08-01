package provider

import (
	"github.com/chapterzero/gomposer/composer"
)

type Provider interface {
	GetApiUrl(string)               string
	GetCommitsUrl(string, string)   string
	GetBranchesUrl(string, string)  string
	GetTagsUrl(string)              string
	GetDownloadUrl(string, string)  string
	GetComposerJson(string, string) composer.ComposerJson
}
