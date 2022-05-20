package database

type ListRequest struct {
	// Empty
}

type CreateRequest struct {
	Name string `json:"name"`
}

type DeleteRequest struct {
	Name string `json:"name"`
}

type UserAddRequest struct {
	Database string `json:"database"`
	Name     string `json:"name"`
}

type UserRemoveRequest struct {
	Database string `json:"database"`
	Name     string `json:"name"`
}

type UserPasswordRequest struct {
	Database string `json:"database"`
	Name     string `json:"name"`
}
