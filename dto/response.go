package dto

type Response struct {
	Meta MetaResponse `json:"meta"`
	Data any          `json:"data"`
}

type MetaResponse struct {
	Code    int    `json:"success"`
	Message string `json:"message"`
	Errors  []any  `json:"errors"`
}
