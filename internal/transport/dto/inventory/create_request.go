package dto

type CreateItemRequest struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

func (r *CreateItemRequest) Validate() error {
	// TODO: Write validation logic here
	return nil
}
