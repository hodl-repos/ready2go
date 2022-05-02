package r2o

type Pagination struct {
	//Offset (Default: 0)
	Offset *int `json:"offset"`

	//Items per page (Default: 25)
	Limit *int `json:"limit"`
}
