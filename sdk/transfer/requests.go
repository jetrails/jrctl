package transfer

type SendRequest struct {
	FilePath string `json:file_path`
}

type ReceiveRequest struct {
	Identifier string `json:"id"`
	Password   string `json:"password"`
}
