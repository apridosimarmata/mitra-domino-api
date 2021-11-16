package Models

type AuthenticationResponse struct {
	Code         int    `json:"code"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
