package request

type CreateAccessoryStateRequest struct {
	AccessoryId uint64 `json:"accessory_id" validate:"required"`
	State       string `json:"state" validate:"required"`
}
