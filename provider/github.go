package provider

const baseUrl = "https://api.github.com/repos";

func GetApiUrl(packageName String) (string) {
	return baseUrl + "/" + packageName
}

func GetCommitsUrl(packageName String, sha String) (string) {
	return GetApiUrl(packageName) + "/commits/" + sha
}

func GetBranchUrl(packageName String, branchName String) (string) {
	return GetApiUrl(packageName) + "/branches/" + branchName
}

func GetTagsUrl(packageName String) (string) {
	return GetApiUrl(packageName) + "/tags"
}
