package website

type ListRequest struct {
	// Empty
}

type SwitchPHPRequest struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
