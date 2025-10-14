package handler

import "fmt"

type CreateMaterialRequest struct {
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	PricePerUnit float64 `json:"pricePerUnit"`
}

type UpdateMaterialRequest struct {
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	PricePerUnit float64 `json:"pricePerUnit"`
}

func (r *CreateMaterialRequest) Validate() error {
	if r.Name == "" || r.Quantity <= 0 || r.PricePerUnit <= 0 {
		return fmt.Errorf("all fields are required and must be valid")
	}

	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}

	if r.Quantity <= 0 {
		return errParamIsRequired("quantity", "int")
	}

	if r.PricePerUnit <= 0 {
		return errParamIsRequired("pricePerUnit", "float")
	}
	return nil
}

func (r *UpdateMaterialRequest) Validate() error {
	if r.Name != "" || r.Quantity >= 0 || r.PricePerUnit >= 0 {
		return nil
	}
	return fmt.Errorf("request body is empty")
}