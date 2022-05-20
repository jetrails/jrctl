package database

type User struct {
	Name string `json:"name"`
	From string `json:"from"`
}

type Database struct {
	Name  string `json:"name"`
	Users []User `json:"users"`
}
