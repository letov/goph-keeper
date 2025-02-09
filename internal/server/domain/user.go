package domain

type User struct {
	BaseEntity
	Email     string
	Pass      string
	PublicKey string
}
