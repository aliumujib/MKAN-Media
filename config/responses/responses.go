package responses

import "fmt"

//GeneralResponse returns a standard response format
type GeneralResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error"`
	Message string      `json:"message"`
}

type InternalError struct {
	Path string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("parse %v: internal error", e.Path)
}
