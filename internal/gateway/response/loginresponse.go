package response

type LoginResponse struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	Expires   string `json:"expires"`
}
