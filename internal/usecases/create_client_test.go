package usecases

import (
	"testing"

	"github.com/jailtonjunior94/go-tests/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientRepositoryMock struct {
	mock.Mock
}

func (c *ClientRepositoryMock) Save(client *entities.Client) error {
	args := c.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	mockRepo := new(ClientRepositoryMock)
	mockRepo.On("Save", mock.Anything).Return(nil)

	useCase := NewCreateClientUseCase(mockRepo)

	input := &CreateClientInputDTO{
		Name:  "Jailton Junior",
		Email: "jailton.junior94@outlook.com",
	}
	output, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, "Jailton Junior", output.Name)
	assert.Equal(t, "jailton.junior94@outlook.com", output.Email)
	assert.Equal(t, 0, output.Points)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "Save", 1)
}
