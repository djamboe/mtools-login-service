package viewmodels

type LoginVM struct {
	Error    bool   `json:"error"`
	Id       int    `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}
