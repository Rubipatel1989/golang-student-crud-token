package request

type StudentSignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
