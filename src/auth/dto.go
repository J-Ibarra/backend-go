package auth

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}