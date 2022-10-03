package error_response

type Error struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Err     error  `json:"error,omitempty"`
}

func (v *Error) Error() string {
	return v.Message
}
