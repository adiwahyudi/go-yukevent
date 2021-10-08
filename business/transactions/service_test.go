package transactions_test

import (
	"testing"
	"yukevent/business"
	"yukevent/business/transactions"
	_transMock "yukevent/business/transactions/mocks"
	randomcode "yukevent/helpers/random-code"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockTransRepository _transMock.Repository
	transService        transactions.Service
)

func TestMain(m *testing.M) {

	transService = transactions.NewServiceTrans(&mockTransRepository)

	m.Run()
}

func TestTrans(t *testing.T) {
	t.Run("test case 1, valid trans", func(t *testing.T) {
		uniq, _ := randomcode.GenerateCode(8)
		outputDomain := transactions.Domain{
			ID:        1,
			UserID:    1,
			EventID:   1,
			Status:    false,
			Uniq_code: uniq,
		}
		inputService := transactions.Domain{
			ID:        1,
			UserID:    1,
			EventID:   1,
			Status:    false,
			Uniq_code: uniq,
		}
		mockTransRepository.On("Trans", mock.Anything, mock.Anything).Return(outputDomain, nil).Once()

		resp, err := transService.Trans(inputService.UserID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.ID, resp.ID)
	})

	t.Run("test case 1, invalid trans", func(t *testing.T) {
		uniq, _ := randomcode.GenerateCode(8)
		outputDomain := transactions.Domain{
			ID:        1,
			UserID:    1,
			EventID:   12,
			Status:    false,
			Uniq_code: uniq,
		}
		inputService := transactions.Domain{
			ID:        1,
			UserID:    11,
			EventID:   1,
			Status:    false,
			Uniq_code: uniq,
		}
		mockTransRepository.On("Trans", mock.Anything, mock.Anything).Return(transactions.Domain{}, business.ErrInternalServer).Once()

		resp, err := transService.Trans(inputService.UserID, &inputService)
		assert.NotNil(t, err)
		assert.NotEqual(t, outputDomain, resp)
	})
}

func TestGetTransByID(t *testing.T) {
	t.Run("test case 1, valid get trans by id", func(t *testing.T) {
		uniq, _ := randomcode.GenerateCode(8)
		expectedReturn := transactions.Domain{
			ID:        1,
			UserID:    1,
			EventID:   12,
			Status:    false,
			Uniq_code: uniq,
		}
		mockTransRepository.On("GetTransByID", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := transService.GetTransByID(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all events", func(t *testing.T) {
		expectedReturn := transactions.Domain{}
		mockTransRepository.On("GetTransByID", mock.Anything).Return(expectedReturn, business.ErrNotFound).Once()
		_, err := transService.GetTransByID(1)
		assert.Equal(t, err, business.ErrNotFound)

	})
}

func TestGetAllTrans(t *testing.T) {
	t.Run("test case 1, valid all trans", func(t *testing.T) {
		uniq, _ := randomcode.GenerateCode(8)
		expectedReturn := []transactions.Domain{
			{
				ID:        1,
				UserID:    1,
				EventID:   12,
				Status:    false,
				Uniq_code: uniq,
			},
			{
				ID:        2,
				UserID:    2,
				EventID:   15,
				Status:    false,
				Uniq_code: uniq,
			},
		}
		mockTransRepository.On("GetAllTrans").Return(expectedReturn, nil).Once()
		_, err := transService.GetAllTrans()
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all trans", func(t *testing.T) {
		expectedReturn := []transactions.Domain{}
		mockTransRepository.On("GetAllTrans").Return(expectedReturn, assert.AnError).Once()
		_, err := transService.GetAllTrans()
		assert.Equal(t, err, assert.AnError)

	})
}

func TestGetAllUserTrans(t *testing.T) {
	t.Run("test case 1, valid all events", func(t *testing.T) {
		uniq, _ := randomcode.GenerateCode(8)
		expectedReturn := []transactions.Domain{
			{
				ID:        1,
				UserID:    2,
				EventID:   12,
				Status:    false,
				Uniq_code: uniq,
			},
			{
				ID:        2,
				UserID:    2,
				EventID:   15,
				Status:    false,
				Uniq_code: uniq,
			},
		}
		mockTransRepository.On("GetAllUserTrans", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := transService.GetAllUserTrans(2)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all events", func(t *testing.T) {
		expectedReturn := []transactions.Domain{}
		mockTransRepository.On("GetAllUserTrans", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := transService.GetAllUserTrans(2)
		assert.Equal(t, err, assert.AnError)

	})
}
