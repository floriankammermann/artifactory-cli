package types

type Repositories struct {
	Repositories []Repository
}

type Repository struct {
	Key				string		`json:"key"`
	Type			string		`json:"type"`
	Url				string		`json:"url"`
}

type RepositoryDetails struct {
	Type        	string 		`json:"rclass"`
	Key         	string 		`json:"key"`
	PackageType 	string 		`json:"packageType"`
	RepoLayoutRef 	string 		`json:"repoLayoutRef"`
}

func CreateRepositoryDetails(repoType string, repoPackageType string, repoKey string, repoLayoutRef string) (details *RepositoryDetails) {
	details = &RepositoryDetails {
		Type:           repoType,
		PackageType:    repoPackageType,
		Key:            repoKey,
		RepoLayoutRef:  repoLayoutRef,
	}
	return details
}

