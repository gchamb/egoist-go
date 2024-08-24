package structs

type GoogleUser struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthResponse struct {
	JwtToken    string `json:"jwt_token"`
	FreshToken  string `json:"refresh_token"`
	IsOnboarded bool   `json:"is_onboarded"`
}
