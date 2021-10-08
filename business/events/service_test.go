package events_test

import (
	"testing"
	"yukevent/business"
	"yukevent/business/events"
	_eventsMock "yukevent/business/events/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockEventRepository _eventsMock.Repository
	eventService        events.Service
)

func TestMain(m *testing.M) {
	// jwtAuth := &middleware.ConfigJWT{
	// 	SecretJWT:       "testauth123",
	// 	ExpiresDuration: 2,
	// }

	eventService = events.NewServiceEvent(&mockEventRepository)

	m.Run()
}

func TestAllEvent(t *testing.T) {
	t.Run("test case 1, valid all events", func(t *testing.T) {
		expectedReturn := []events.Domain{
			{
				ID:               1,
				OrganizerID:      1,
				Name:             "ABC",
				Description:      "ABC",
				Syarat_Ketentuan: "ABC",
				Location:         "ABC",
				Event_Date:       "ABC",
				Event_Time:       "ABC",
				Close_Register:   "ABC",
				Capacity:         1,
				Poster:           "ABC",
				Price:            1,
			},
			{
				ID:               2,
				OrganizerID:      1,
				Name:             "ABCABC",
				Description:      "ABCABC",
				Syarat_Ketentuan: "ABCABC",
				Location:         "ABCABC",
				Event_Date:       "ABCABC",
				Event_Time:       "ABCABC",
				Close_Register:   "ABCABC",
				Capacity:         10,
				Poster:           "ABCABC",
				Price:            10,
			},
		}
		mockEventRepository.On("AllEvent").Return(expectedReturn, nil).Once()
		_, err := eventService.AllEvent()
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all events", func(t *testing.T) {
		expectedReturn := []events.Domain{}
		mockEventRepository.On("AllEvent").Return(expectedReturn, assert.AnError).Once()
		_, err := eventService.AllEvent()
		assert.Equal(t, err, assert.AnError)

	})
}

func TestCreate(t *testing.T) {
	t.Run("test case 1, valid create event", func(t *testing.T) {
		outputDomain := events.Domain{
			ID:               2,
			OrganizerID:      1,
			Name:             "ABCABC",
			Description:      "ABCABC",
			Syarat_Ketentuan: "ABCABC",
			Location:         "ABCABC",
			Event_Date:       "ABCABC",
			Event_Time:       "ABCABC",
			Close_Register:   "ABCABC",
			Capacity:         10,
			Poster:           "ABCABC",
			Price:            10,
		}
		inputService := events.Domain{
			ID:               2,
			OrganizerID:      1,
			Name:             "ABCABC",
			Description:      "ABCABC",
			Syarat_Ketentuan: "ABCABC",
			Location:         "ABCABC",
			Event_Date:       "ABCABC",
			Event_Time:       "ABCABC",
			Close_Register:   "ABCABC",
			Capacity:         10,
			Poster:           "ABCABC",
			Price:            10,
		}
		mockEventRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := eventService.Create(inputService.OrganizerID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Name, resp.Name)
	})

	t.Run("test case 2, invalid create event", func(t *testing.T) {
		outputDomain := events.Domain{}
		inputService := events.Domain{
			ID:               2,
			OrganizerID:      1,
			Name:             "ABCABC",
			Description:      "ABCABC",
			Syarat_Ketentuan: "ABCABC",
			Location:         "ABCABC",
			Event_Date:       "ABCABC",
			Event_Time:       "ABCABC",
			Close_Register:   "ABCABC",
			Capacity:         10,
			Poster:           "ABCABC",
			Price:            10,
		}
		mockEventRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()

		resp, err := eventService.Create(inputService.OrganizerID, &inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid update event", func(t *testing.T) {
		outputDomain := events.Domain{
			ID:               2,
			OrganizerID:      1,
			Name:             "ABCABC",
			Description:      "ABCABC",
			Syarat_Ketentuan: "ABCABC",
			Location:         "ABCABC",
			Event_Date:       "ABCABC",
			Event_Time:       "ABCABC",
			Close_Register:   "ABCABC",
			Capacity:         10,
			Poster:           "ABCABC",
			Price:            10,
		}
		inputService := events.Domain{
			ID:               2,
			OrganizerID:      1,
			Name:             "ABCABC",
			Description:      "ABCABC",
			Syarat_Ketentuan: "ABCABC",
			Location:         "ABCABC",
			Event_Date:       "ABCABC",
			Event_Time:       "ABCABC",
			Close_Register:   "ABCABC",
			Capacity:         10,
			Poster:           "ABCABC",
			Price:            10,
		}
		mockEventRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := eventService.Update(inputService.OrganizerID, inputService.ID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Name, resp.Name)
	})

	t.Run("test case 2, invalid update event", func(t *testing.T) {
		outputDomain := events.Domain{}
		inputService := events.Domain{
			ID:               2,
			OrganizerID:      1,
			Name:             "ABCABC",
			Description:      "ABCABC",
			Syarat_Ketentuan: "ABCABC",
			Location:         "ABCABC",
			Event_Date:       "ABCABC",
			Event_Time:       "ABCABC",
			Close_Register:   "ABCABC",
			Capacity:         10,
			Poster:           "ABCABC",
			Price:            10,
		}
		mockEventRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()
		resp, err := eventService.Update(inputService.OrganizerID, inputService.ID, &inputService)

		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid delete event", func(t *testing.T) {
		mockEventRepository.On("Delete", mock.Anything, mock.Anything).Return("Event has been delete", nil).Once()
		resp, err := eventService.Delete(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, "Event has been delete", resp)
	})

	t.Run("test case 2, invalid delete event", func(t *testing.T) {
		mockEventRepository.On("Delete", mock.Anything, mock.Anything).Return("", business.ErrNotFound).Once()
		resp, err := eventService.Delete(2, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestMyEventByOrganizer(t *testing.T) {
	t.Run("test case 1, valid all events", func(t *testing.T) {
		expectedReturn := []events.Domain{
			{
				ID:               1,
				OrganizerID:      1,
				Name:             "ABC",
				Description:      "ABC",
				Syarat_Ketentuan: "ABC",
				Location:         "ABC",
				Event_Date:       "ABC",
				Event_Time:       "ABC",
				Close_Register:   "ABC",
				Capacity:         1,
				Poster:           "ABC",
				Price:            1,
			},
			{
				ID:               2,
				OrganizerID:      1,
				Name:             "ABCABC",
				Description:      "ABCABC",
				Syarat_Ketentuan: "ABCABC",
				Location:         "ABCABC",
				Event_Date:       "ABCABC",
				Event_Time:       "ABCABC",
				Close_Register:   "ABCABC",
				Capacity:         10,
				Poster:           "ABCABC",
				Price:            10,
			},
		}
		mockEventRepository.On("MyEventByOrganizer", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := eventService.MyEventByOrganizer(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all events", func(t *testing.T) {
		expectedReturn := []events.Domain{}
		mockEventRepository.On("MyEventByOrganizer", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := eventService.MyEventByOrganizer(1)
		assert.Equal(t, err, assert.AnError)

	})
}

func TestEventByID(t *testing.T) {
	t.Run("test case 1, valid all events", func(t *testing.T) {
		expectedReturn := events.Domain{
			ID:               1,
			OrganizerID:      1,
			Name:             "ABC",
			Description:      "ABC",
			Syarat_Ketentuan: "ABC",
			Location:         "ABC",
			Event_Date:       "ABC",
			Event_Time:       "ABC",
			Close_Register:   "ABC",
			Capacity:         1,
			Poster:           "ABC",
			Price:            1,
		}
		mockEventRepository.On("EventByID", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := eventService.EventByID(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all events", func(t *testing.T) {
		expectedReturn := events.Domain{}
		mockEventRepository.On("EventByID", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := eventService.EventByID(1)
		assert.Equal(t, err, assert.AnError)

	})
}

func TestEventByIdOrganizer(t *testing.T) {
	t.Run("test case 1, valid all events", func(t *testing.T) {
		expectedReturn := []events.Domain{
			{
				ID:               1,
				OrganizerID:      1,
				Name:             "ABC",
				Description:      "ABC",
				Syarat_Ketentuan: "ABC",
				Location:         "ABC",
				Event_Date:       "ABC",
				Event_Time:       "ABC",
				Close_Register:   "ABC",
				Capacity:         1,
				Poster:           "ABC",
				Price:            1,
			},
			{
				ID:               2,
				OrganizerID:      1,
				Name:             "ABCABC",
				Description:      "ABCABC",
				Syarat_Ketentuan: "ABCABC",
				Location:         "ABCABC",
				Event_Date:       "ABCABC",
				Event_Time:       "ABCABC",
				Close_Register:   "ABCABC",
				Capacity:         10,
				Poster:           "ABCABC",
				Price:            10,
			},
		}
		mockEventRepository.On("EventByIdOrganizer", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := eventService.EventByIdOrganizer(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all events", func(t *testing.T) {
		expectedReturn := []events.Domain{}
		mockEventRepository.On("EventByIdOrganizer", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := eventService.EventByIdOrganizer(1)
		assert.Equal(t, err, assert.AnError)

	})
}
