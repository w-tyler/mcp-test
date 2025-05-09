package entity

type Product struct {
	ID          string
	Name        string
	Description string
	Price       Money // Value object
	Stock       int
}
