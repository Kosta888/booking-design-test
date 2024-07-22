package domain

import (
    "errors"
    "time"
)

var (
    ErrRoomNotAvailable = errors.New("room is not available")
)

type Order struct {
    ID        int       `json:"id,omitempty"`
    HotelID   string    `json:"hotel_id"`
    RoomID    string    `json:"room_id"`
    UserEmail string    `json:"email"`
    From      time.Time `json:"from"`
    To        time.Time `json:"to"`
}

type RoomAvailability struct {
    HotelID string    `json:"hotel_id"`
    RoomID  string    `json:"room_id"`
    Date    time.Time `json:"date"`
    Quota   int       `json:"quota"`
}
