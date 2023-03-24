package alternative

type Entry struct {
	Name       string `json:"name"`
	Current string   `json:"current"`
	Versions []string `json:"versions"`
}