package server

type TokenData struct {
	Identity         string   `json:"identity"`
	TokenID          string   `json:"token_id"`
	AllowedClientIPs []string `json:"allowed_client_ips"`
}
