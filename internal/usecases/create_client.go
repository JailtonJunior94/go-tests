package usecases

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID     string
	Name   string
	Email  string
	Points int
}
