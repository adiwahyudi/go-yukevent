package organizers_test

import (
	"testing"
	"yukevent/app/middleware"
	"yukevent/business"
	"yukevent/business/organizers"
	_organziersMock "yukevent/business/organizers/mocks"
	"yukevent/helpers/encrypt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockOrganizerRepository _organziersMock.Repository
	organizersnService      organizers.Service
)

func TestMain(m *testing.M) {
	jwtAuth := &middleware.ConfigJWT{
		SecretJWT:       "testauth123",
		ExpiresDuration: 2,
	}

	organizersnService = organizers.NewServiceOrganizer(&mockOrganizerRepository, 2, jwtAuth)

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("123456")
		outputDomain := organizers.Domain{
			Username:     "iwayanadiwahyudi",
			Email:        "iwayanadiwahyudi@mail.com",
			Password:     password,
			Name:         "I Wayan Adi Wahyudi",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}
		inputService := organizers.Domain{
			Username:     "iwayanadiwahyudi",
			Email:        "iwayanadiwahyudi@mail.com",
			Password:     "123456",
			Name:         "I Wayan Adi Wahyudi",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}
		mockOrganizerRepository.On("Register", mock.Anything).Return(outputDomain, nil).Once()

		resp, err := organizersnService.Register(&inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Username, resp.Username)
	})

	t.Run("test case 2, invalid register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("654321")
		outputDomain := organizers.Domain{
			Username:     "iwayanadiwahyudi",
			Email:        "iwayanadiwahyudi@mail.com",
			Password:     password,
			Name:         "I Wayan Adi Wahyudi",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}
		inputService := organizers.Domain{
			Username:     "iwayanadiwahyudi",
			Email:        "iwayanadiwahyudi@mail.com",
			Password:     "123456",
			Name:         "I Wayan Adi Wahyudi",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}
		mockOrganizerRepository.On("Register", mock.Anything).Return(outputDomain, business.ErrInternalServer).Once()

		resp, err := organizersnService.Register(&inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("test case 3, invalid register duplicate", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("123456")
		outputDomain := organizers.Domain{
			Username:     "iwayanadiwahyudi",
			Email:        "iwayanadiwahyudi@mail.com",
			Password:     password,
			Name:         "I Wayan Adi Wahyudi",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}
		inputService := organizers.Domain{
			Username:     "iwayanadiwahyudi",
			Email:        "iwayanadiwahyudi@mail.com",
			Password:     "123456",
			Name:         "I Wayan Adi Wahyudi",
			Phone_Number: "0811111111",
			Photo:        "ini_poto",
		}

		mockOrganizerRepository.On("Register", mock.Anything).Return(organizers.Domain{}, business.ErrDuplicateData).Once()

		resp, err := organizersnService.Register(&inputService)

		assert.NotNil(t, err)
		assert.NotEqual(t, outputDomain, resp)
	})
}

func TestLogin(t *testing.T) {
	t.Run("test case 1, valid login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("123456")
		outputDomain := organizers.Domain{
			Email:    "iwayanadiwahyudi@mail.com",
			Password: password,
		}
		inputService := organizers.Domain{
			Email:    "iwayanadiwahyudi@mail.com",
			Password: "123456",
		}
		mockOrganizerRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, nil).Once()

		resp, err := organizersnService.Login(inputService.Email, inputService.Password)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("test case 2, invalid login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("654321")
		outputDomain := organizers.Domain{
			Email:    "iwayanadiwahyudi@mail.com",
			Password: password,
		}
		inputService := organizers.Domain{
			Email:    "iwayanadiwahyudi@mail.com",
			Password: "123456",
		}
		mockOrganizerRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, business.ErrEmailorPass).Once()

		resp, err := organizersnService.Login(inputService.Email, inputService.Password)
		assert.Empty(t, resp)
		assert.Equal(t, err, business.ErrEmailorPass)
	})

	t.Run("test case 4, invalid login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("6127376")
		outputDomain := organizers.Domain{
			Email:    "iwayanadiwahyudi@mail.com",
			Password: password,
		}
		inputService := organizers.Domain{
			Email:    "iwayanadiwahyudizzz@mail.com",
			Password: "123456",
		}
		mockOrganizerRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, business.ErrEmailorPass).Once()

		resp, err := organizersnService.Login(inputService.Email, inputService.Password)
		assert.Empty(t, resp)
		assert.Equal(t, err, business.ErrEmailorPass)
	})
}
