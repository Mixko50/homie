package response

import "time"

type GetAccessoryResponse struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	GroupId   uint64    `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
