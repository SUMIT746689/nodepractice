package domain

type Company struct {
	Name   string `json:"name" validate:"required,max=50"`
	Domain string `json:"domain" validate:"required,max=50"`
}

type UpdateCompanyRequest struct {
	Name   string `json:"name" validate:"max=50"`
	Domain string `json:"domain" validate:"max=50"`
}

type Representative struct {
	Name  string
	Phone string
}

type RepresentativeJson struct {
	Name  string `json:"name" validate:"max=50"`
	Phone string `json:"phone" validate:"max=50"`
}
