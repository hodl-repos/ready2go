package r2o

type Pagination struct {
	//Page (Default: 1)
	Page *int `url:"page"`

	//Items per page (Default: 25)
	Limit *int `url:"limit"`
}
