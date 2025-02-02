package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientinfo"`
	// .. alot more
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
