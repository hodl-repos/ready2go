package r2o

type Ready2GoService struct {
	developerApiToken *string
}

func InitializeNewService(developerApiToken string) Ready2GoService {
	return Ready2GoService{
		developerApiToken: &developerApiToken,
	}
}
