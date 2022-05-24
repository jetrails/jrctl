package database

type User struct {
	Name string `json:"name"`
	From string `json:"from"`
}

type UserWithDatabases struct {
	User
	Databases []Database `json:"databases"`
}

type Database struct {
	Name string `json:"name"`
}

type DatabaseWithUsers struct {
	Database
	Users []User `json:"users"`
}

type ConfirmPayload struct {
	Identifier string `json:"id"`
	TTL        string `json:"ttl"`
}
