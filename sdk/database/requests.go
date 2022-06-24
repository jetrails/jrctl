package database

type ListDatabasesRequest struct {
	// Empty
}

type ListUsersRequest struct {
	// Empty
}

type CreateRequest struct {
	Name string `json:"name"`
}

type DeleteRequest struct {
	Name string `json:"name"`
}

type UserCreateRequest struct {
	Name   string `json:"name"`
	From   string `json:"from"`
	Native bool   `json:"native"`
}

type UserDeleteRequest struct {
	Name string `json:"name"`
	From string `json:"from"`
}

type UserPasswordRequest struct {
	Name   string `json:"name"`
	From   string `json:"from"`
	Native bool   `json:"native"`
}

type LinkRequest struct {
	Database string `json:"database"`
	Name     string `json:"name"`
	From     string `json:"from"`
}

type UnlinkRequest struct {
	Database string `json:"database"`
	Name     string `json:"name"`
	From     string `json:"from"`
}
