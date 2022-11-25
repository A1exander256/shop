package models

type Product struct {
	Id          int            `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Description string         `json:"description" db:"description"`
	LeftInStock uint           `json:"left_in_stock " db:"left_in_stock" binding:"min=0"`
	Prices      []PriceProduct `json:"prices" db:"prices"`
}

type PriceProduct struct {
	ProductId int     `json:"product_id" db:"product_id"`
	Currency  string  `json:"currency" db:"currency"`
	Price     float64 `json:"price" db:"price"`
}
