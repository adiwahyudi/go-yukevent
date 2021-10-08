package admins_test

import (
	"testing"
	"yukevent/app/middleware"
	"yukevent/business/admins"
	_adminMock "yukevent/business/admins/mocks"
	"yukevent/business/users"
	"yukevent/helpers/encrypt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockAdminRepository _adminMock.Repository
	adminService        admins.Service
)

func TestMain(m *testing.M) {
	jwtAuth := &middleware.ConfigJWT{
		SecretJWT:       "testauth123",
		ExpiresDuration: 2,
	}

	adminService = admins.NewServiceAdmin(&mockAdminRepository, 2, jwtAuth)

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid test for register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("123456")
		outputDomain := users.Domain{
			ID:       1,
			Username: "adiwahyudi",
			Password: "123456",
		}
		inputService := users.Domain{
			ID:       1,
			Username: "iwayanadiwahyudi",
			Password: "123456",
		}
		mockAdminRepository.On("Register", mock.Anything).Return(outputDomain, nil).Once()

		resp, err := adminService.Register(&inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Username, resp.Username)
	})
}
