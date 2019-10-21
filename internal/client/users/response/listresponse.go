package response

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type ListResponse struct {
	Users []User `json:"users"`
}
