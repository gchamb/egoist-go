package structs

type GoogleUser struct {
	Email   string `json: "email"`
	Name    string `json: "name"`
	Picture string `json: "picture"`
}