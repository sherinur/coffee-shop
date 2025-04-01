package model

type Orders struct {
	ID           int
	CustomerName string
	Status       string
	Notes        string
	CreatedAt    string
}

func (r *Orders) Validate() error {
	switch {
	case r.ID <= 0:
		return ErrNotValidOrderID
	case r.CustomerName == "":
		return ErrNotValidOrderCustomerName
	case r.Status == "":
		return ErrNotValidOrderStatus
	case r.CreatedAt == "":
		return ErrNotValidCreatedAt
	default:
		return nil
	}
}
