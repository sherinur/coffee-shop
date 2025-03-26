package models

type InventoryTransactions struct {
	TransactionID  int
	IngredientId   int
	QuantityChange float64
	Reason         string
	CreatedAt      string
}
