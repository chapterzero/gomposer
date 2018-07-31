package composer

type ComposerJson struct {
	Name			string
	Require			map[string]string
	Repositories	[]Repository
}

type Repository struct {
	Type			string
	Url				string
}
