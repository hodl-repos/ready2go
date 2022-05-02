package r2o

type Pagination struct {
	//Offset (Default: 0)
	Offset *int `url:"offset"`

	//Items per page (Default: 25)
	Limit *int `url:"limit"`
}
