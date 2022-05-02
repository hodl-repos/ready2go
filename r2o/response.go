package r2o

type MessageResponse struct {
	Error   *bool   `json:"error"`
	Success *bool   `json:"success"`
	Message *string `json:"msg"`
}
