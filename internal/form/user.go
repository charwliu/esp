package form

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Enabled  bool   `json:"enabled"`
	Gravatar bool   `json:"gravatar"`
}
