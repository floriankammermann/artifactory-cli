package types

type Repositories struct {
	Repositories []Repository
}

type Repository struct {
	Key				string		`json:"key"`
	Type			string		`json:"type"`
	Url				string		`json:"url"`
}
