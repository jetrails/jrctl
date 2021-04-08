package firewall

type AllowRequest struct {
	Address  string `json:"address"`
	Ports    []int  `json:"ports"`
	Protocol string `json:"protocol"`
	Comment  string `json:"comment"`
}

type UnAllowRequest struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type DenyRequest struct {
	Address  string `json:"address"`
	Ports    []int  `json:"ports"`
	Protocol string `json:"protocol"`
	Comment  string `json:"comment"`
}

type UnDenyRequest struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type AllowCloudflareRequest struct {
	// Empty
}

type UnAllowCloudflareRequest struct {
	// Empty
}
