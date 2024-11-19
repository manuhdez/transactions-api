package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/test/mocks"
)

func TestUsersRetriever_Retrieve(t *testing.T) {
	testUsers := []user.User{
		{Id: "23", FirstName: "Mark", LastName: "Zuckerberg", Email: "mark@zuckerberg.com", Password: "1234"},
	}

	t.Run("should return a list of users", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		repo.On("All", mock.Anything).Return(testUsers, nil)

		retriever := NewUsersRetrieverService(repo)
		got, err := retriever.Retrieve()
		require.NoError(t, err)

		assert.Equal(t, got, testUsers, "expected users to be equal")
	})

	t.Run("should return an error if the repository fails", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		repo.On("All", mock.Anything).Return(nil, fmt.Errorf("failed to retrieve users"))

		retriever := NewUsersRetrieverService(repo)
		_, err := retriever.Retrieve()
		assert.EqualError(t, err, "failed to retrieve users", "expected an error")
	})
}
