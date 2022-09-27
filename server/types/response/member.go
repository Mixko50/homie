package response

import "time"

type GetMemberResponse struct {
	Id         uint64    `json:"id"`
	Name       string    `json:"name"`
	DeviceName string    `json:"device_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateMemberResponse struct {
	Token string `json:"token"`
}
