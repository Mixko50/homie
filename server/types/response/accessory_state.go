package response

import "time"

type GetAccessoryStateResponse struct {
	Id          uint64    `json:"id"`
	AccessoryId uint64    `json:"accessory_id"`
	State       string    `json:"state"`
	GroupId     uint64    `json:"group_id"`
	GroupName   string    `json:"group_name"`
	MemberId    uint64    `json:"member_id"`
	MemberName  string    `json:"member_name"`
	CreatedAt   time.Time `json:"created_at"`
}
