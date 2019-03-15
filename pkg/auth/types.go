package auth

// RegisterRequest is the JSON payload format for /register
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthorizeRequest is the JSON payload format for /auth
type AuthorizeRequest struct {
	Token string `json:"token,omitempty"`
}

// LogoutRequest is the JSON payload format for /logout
type LogoutRequest struct {
}
