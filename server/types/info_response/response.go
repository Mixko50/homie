package info_response

type InfoResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Info    string `json:"info,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewInfoResponse(args1 any, args2 ...any) *InfoResponse {
	if msg, ok := args1.(string); ok {
		if len(args2) != 0 {
			return &InfoResponse{
				Success: true,
				Code:    msg,
				Info:    args2[0].(string),
			}
		}
		return &InfoResponse{
			Success: true,
			Info:    msg,
		}
	}

	return &InfoResponse{
		Success: true,
		Code:    "DATA",
		Data:    args1,
	}
}
