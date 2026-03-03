package request

type UpdateUserRequest struct {
	FirstName   *string `json:"first_name" form:"first_name" validate:"omitempty,min=2,max=50"`
	LastName    *string `json:"last_name" form:"last_name" validate:"omitempty,min=2,max=50"`
	Username    *string `json:"username" form:"username" validate:"omitempty,min=3,max=30"`
	PhoneNumber *string `json:"phone_number" form:"phone_number" validate:"omitempty,e164"`
}
