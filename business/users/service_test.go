package users_test

import (
	"testing"
	"yukevent/app/middleware"
	"yukevent/business/users"
	_userMock "yukevent/business/users/mocks"
	"yukevent/helpers/encrypt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockUserRepository _userMock.Repository
	usersService       users.Service
)

func TestMain(m *testing.M) {
	jwtAuth := &middleware.ConfigJWT{
		SecretJWT:       "testauth123",
		ExpiresDuration: 2,
	}

	usersService = users.NewServiceUser(&mockUserRepository, 2, jwtAuth)

	m.Run()
}

func TestLogin(t *testing.T) {
	t.Run("test case 1 | valid test", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("adi123")

		outputDomain := users.Domain{
			Email:    "adi@email.com",
			Password: password,
		}

		inputService := users.Domain{
			Email:    "adi@email.com",
			Password: "adi123",
		}

		mockUserRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, nil).Once()

		resp, err := usersService.Login(inputService.Email, inputService.Password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("test case 2 | wrong password test", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("adi123")
		outputDomain := users.Domain{
			Email:    "adi@email.com",
			Password: password,
		}

		inputService := users.Domain{
			Email:    "adi@email.com",
			Password: "adi321",
		}

		mockUserRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, assert.AnError).Once()

		resp, err := usersService.Login(inputService.Email, inputService.Password)
		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})

	t.Run("test case 3 | no email and password test", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("adi123")
		outputDomain := users.Domain{
			Email:    "adi@email.com",
			Password: password,
		}

		inputService := users.Domain{
			Email:    "adi321@email.com",
			Password: "adi321",
		}

		mockUserRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, assert.AnError).Once()

		resp, err := usersService.Login(inputService.Email, inputService.Password)
		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})
}
