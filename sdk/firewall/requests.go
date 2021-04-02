package firewall

type AllowRequest struct {
	Address string  `json:"address"`
	Ports [] int    `json:"ports"`
	Blame string    `json:"blame"`
	Comment string  `json:"comment"`
}

type CloudflareRequest struct {
	Blame string    `json:"blame"`
}
