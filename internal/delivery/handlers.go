package delivery

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/logger"
	"applicationDesignTest/internal/usecase"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func CreateOrderHandler(u usecase.BookingUsecase, log logger.Interface, validate *validator.Validate) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var order domain.Order
        if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
            log.Error(err)
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        if err := validate.Struct(order); err != nil {
            log.Error(err)
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        if err := u.CreateOrder(&order); err != nil {
            log.Error(err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        log.Info("Order created successfully")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(order)
    }
}
