package website

type Properties struct {
	Name       string `json:"name"`
	Compatible bool   `json:"compatible"`
	PHPVersion string `json:"php_version"`
}

type Availablity struct {
	Name       string `json:"name"`
	Installed  bool   `json:"installed"`
	Enabled    bool   `json:"enabled"`
	Configured bool   `json:"configured"`
}
