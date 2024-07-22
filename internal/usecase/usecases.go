package usecase

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/logger"
	"applicationDesignTest/internal/memory"
	"time"
)

type BookingUsecase interface {
    CreateOrder(order *domain.Order) error
}

type bookingUsecase struct {
    orderRepo        *memory.InMemoryOrderRepository
    availabilityRepo *memory.InMemoryRoomAvailabilityRepository
    logger           logger.Interface
}

func NewBookingUsecase(orderRepo *memory.InMemoryOrderRepository, availabilityRepo *memory.InMemoryRoomAvailabilityRepository, log logger.Interface) BookingUsecase {
    return &bookingUsecase{
        orderRepo:        orderRepo,
        availabilityRepo: availabilityRepo,
        logger:           log,
    }
}

func (u *bookingUsecase) CreateOrder(order *domain.Order) error {
    // Проверка доступности номера
    available, err := u.isRoomAvailable(order.HotelID, order.RoomID, order.From, order.To)
    if err != nil {
        return err
    }
    if !available {
        return domain.ErrRoomNotAvailable
    }

    // Создание заказа
    if err := u.orderRepo.Create(order); err != nil {
        u.logger.Error(err.Error())
        return err
    }

    // Обновление доступности номера
    if err := u.updateRoomAvailability(order.HotelID, order.RoomID, order.From, order.To); err != nil {
        return err
    }

    u.logger.Info("Order created successfully with ID %d", order.ID)
    return nil
}

func (u *bookingUsecase) isRoomAvailable(hotelID, roomID string, from, to time.Time) (bool, error) {
    for date := from; !date.After(to); date = date.AddDate(0, 0, 1) {
        availability, err := u.availabilityRepo.GetByHotelIDAndRoomIDAndDate(hotelID, roomID, date)
        if err != nil {
            return false, err
        }
        if availability.Quota <= 0 {
            return false, nil
        }
    }
    return true, nil
}

func (u *bookingUsecase) updateRoomAvailability(hotelID, roomID string, from, to time.Time) error {
    for date := from; !date.After(to); date = date.AddDate(0, 0, 1) {
        if err := u.availabilityRepo.DecreaseQuota(hotelID, roomID, date); err != nil {
            return err
        }
    }
    return nil
}
