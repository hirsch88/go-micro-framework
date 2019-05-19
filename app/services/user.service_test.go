package services

import (
	"github.com/hirsch88/go-micro-framework/app/models"
	"github.com/hirsch88/go-micro-framework/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func testSetup(t *testing.T) (UserService, models.User) {
	var user = models.User{
		Username: "bubu",
		Email:    "bubu@test.ch",
	}

	var mailProviderMock = new(mocks.MailProvider)
	var userRepositoryMock = new(mocks.UserRepository)

	userRepositoryMock.On("Create", mock.Anything).Return(user)
	mailProviderMock.On("Send", mock.Anything, mock.Anything, mock.Anything)

	var service = NewUserService(userRepositoryMock, mailProviderMock)
	return service, user
}

func TestCreate(t *testing.T) {
	service, user := testSetup(t)

	returnedUser := service.Create(user)

	assert.Equal(t, "bubu", returnedUser.Username)
	assert.Equal(t, "bubu@test.ch", returnedUser.Email)
}
