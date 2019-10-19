package response

type CheckPasswordResponse struct {
	ID       string `json: "id"`
	Password string `json: "password"`
}
