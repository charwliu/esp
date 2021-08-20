package form

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Enabled  bool   `json:"enabled"`
	Password string `json:"password"`
	Gravatar bool   `json:"gravatar"`
	Mobile   string `json:"mobile"`
}
