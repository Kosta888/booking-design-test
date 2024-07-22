package app

import (
	"applicationDesignTest/internal/delivery"
	"applicationDesignTest/internal/logger"
	"applicationDesignTest/internal/memory"
	"applicationDesignTest/internal/populate"
	"applicationDesignTest/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

type App struct {
    Router                 *chi.Mux
    BookingUsecase         usecase.BookingUsecase
    OrderRepository        *memory.InMemoryOrderRepository
    AvailabilityRepository *memory.InMemoryRoomAvailabilityRepository
    Logger                 logger.Interface
}

func NewApp(logLevel string) *App {
    log := logger.New(logLevel)
    orderRepo := memory.NewInMemoryOrderRepository()
    availabilityRepo := memory.NewInMemoryRoomAvailabilityRepository()

    populate.PopulateInitialData(availabilityRepo)

    bookingUsecase := usecase.NewBookingUsecase(orderRepo, availabilityRepo, log)

    router := chi.NewRouter()
    router.Use(middleware.Logger)

    validate := validator.New()
    router.Post("/orders", delivery.CreateOrderHandler(bookingUsecase, log, validate))

    log.Info("Application initialized")

    return &App{
        Router:                 router,
        BookingUsecase:         bookingUsecase,
        OrderRepository:        orderRepo,
        AvailabilityRepository: availabilityRepo,
        Logger:                 log,
    }
}
