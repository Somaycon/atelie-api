package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func (r *CreateProductRequest) Validate() error {

	if r.Name == "" && r.Price == 0 && r.Description == "" {
		return fmt.Errorf("request body is empty")
	}

	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Price <= 0 {
		return errParamIsRequired("price", "float")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	return nil
}