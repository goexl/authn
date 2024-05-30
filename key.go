package authn

type Key struct {
	Id         string `json:"id,omitempty" validate:"required"`
	Credential string `json:"credential,omitempty" validate:"required"`
	Type       Type   `json:"type,omitempty" validate:"omitempty,oneof=username email qq wechat mobile"`
}
