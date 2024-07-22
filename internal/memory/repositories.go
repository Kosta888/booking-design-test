package memory

import (
	"applicationDesignTest/internal/domain"
	"errors"
	"sync"
	"time"
)

var (
    ErrRoomNotFound = errors.New("room not found")
)

type InMemoryOrderRepository struct {
    mu     sync.RWMutex
    orders []domain.Order
    nextID int
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
    return &InMemoryOrderRepository{
        orders: []domain.Order{},
        nextID: 1,
    }
}

func (r *InMemoryOrderRepository) Create(order *domain.Order) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    order.ID = r.nextID
    r.nextID++
    r.orders = append(r.orders, *order)
    return nil
}

type InMemoryRoomAvailabilityRepository struct {
    mu            sync.RWMutex
    availabilities []domain.RoomAvailability
}

func NewInMemoryRoomAvailabilityRepository() *InMemoryRoomAvailabilityRepository {
    return &InMemoryRoomAvailabilityRepository{
        availabilities: []domain.RoomAvailability{},
    }
}

func (r *InMemoryRoomAvailabilityRepository) Create(availability *domain.RoomAvailability) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.availabilities = append(r.availabilities, *availability)
}

func (r *InMemoryRoomAvailabilityRepository) GetByHotelIDAndRoomIDAndDate(hotelID, roomID string, date time.Time) (domain.RoomAvailability, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    for _, availability := range r.availabilities {
        if availability.HotelID == hotelID && availability.RoomID == roomID && availability.Date.Equal(date) {
            return availability, nil
        }
    }

    return domain.RoomAvailability{}, ErrRoomNotFound
}

func (r *InMemoryRoomAvailabilityRepository) DecreaseQuota(hotelID, roomID string, date time.Time) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    for i, availability := range r.availabilities {
        if availability.HotelID == hotelID && availability.RoomID == roomID && availability.Date.Equal(date) {
            if availability.Quota > 0 {
                r.availabilities[i].Quota--
                return nil
            }
            return errors.New("no quota available")
        }
    }

    return ErrRoomNotFound
}
