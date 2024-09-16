package domain

type RequestBmi struct {
	High   float64 `json:"high" validate:"required"`
	Weight float64 `json:"weight" validate:"required"`
}
