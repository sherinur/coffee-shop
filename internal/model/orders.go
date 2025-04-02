package model

type Order struct {
	ID           int
	CustomerName string
	Status       string
	Notes        string
}

// TODO: Write inventory suffiency validation
// TODO: Implement status and date logic

func (r *Order) Validate() error {
	switch {
	case r.ID <= 0:
		return ErrNotValidOrderID
	case r.CustomerName == "":
		return ErrNotValidOrderCustomerName
	case r.Status == "":
		return ErrNotValidOrderStatus
	default:
		return nil
	}
}
