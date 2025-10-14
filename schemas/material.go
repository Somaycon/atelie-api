package schemas

type Material struct {
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	PricePerUnit float64 `json:"pricePerUnit"`
}