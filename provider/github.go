package provider

const baseUrl = "https://api.github.com/repos";

type Github struct {}

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
