package entity

type User struct {
	ID       string
	Name     string
	Email    string
	Address  Address // Value object
}
