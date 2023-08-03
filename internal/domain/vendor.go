package domain

type Vendor struct {
	Name           string              `json:"name" validate:"required,max=50"`
	Address        string              `json:"address" validate:"required,max=50"`
	Email          string              `json:"email" validate:"omitempty,email"`
	Representative *RepresentativeJson `json:"representative" validate:"required"`
}

type UpdateVendor struct {
	Name           string              `json:"name" validate:"max=50"`
	Address        string              `json:"address" validate:"max=50"`
	Email          string              `json:"email" validate:"omitempty,email"`
	Representative *RepresentativeJson `json:"representative" validate:""`
}
