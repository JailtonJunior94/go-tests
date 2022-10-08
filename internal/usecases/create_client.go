package usecases

import "github.com/jailtonjunior94/go-tests/internal/entities"

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

type CreateClientUseCase struct {
	ClientRepository entities.ClientRepositoryInterface
}

func NewCreateClientUseCase(clientRepository entities.ClientRepositoryInterface) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientRepository: clientRepository,
	}
}

func (c *CreateClientUseCase) Execute(input *CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entities.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = c.ClientRepository.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:     client.ID,
		Name:   client.Name,
		Email:  client.Email,
		Points: client.Points,
	}, nil

}
