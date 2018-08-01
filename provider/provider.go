package provider

type Provider interface {
	GetApiUrl(string)               string
	GetCommitsUrl(string, string)   string
	GetBranchesUrl(string, string)  string
	GetTagsUrl(string)              string
}
