package dto

type SaveUser struct {
	Email     string
	PassHash  string
	PublicKey string
}

type LoginUser struct {
	Email    string
	PassHash string
}
