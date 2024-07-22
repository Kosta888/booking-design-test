package populate

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/memory"
	"time"
)

var initialData = []domain.RoomAvailability{
	{"reddison", "lux", date(2024, 1, 1), 1},
	{"reddison", "lux", date(2024, 1, 2), 1},
	{"reddison", "lux", date(2024, 1, 3), 1},
	{"reddison", "lux", date(2024, 1, 4), 1},
	{"reddison", "lux", date(2024, 1, 5), 0},
	{"reddison", "lux", date(2024, 1, 6), 1},
	{"reddison", "lux", date(2024, 1, 7), 1},
	{"reddison", "lux", date(2024, 1, 8), 1},
	{"reddison", "lux", date(2024, 1, 9), 1},
	{"reddison", "lux", date(2024, 1, 10), 0},
}

func PopulateInitialData(repo *memory.InMemoryRoomAvailabilityRepository) {

	for _, availability := range initialData {
		repo.Create(&availability)
	}
}

func date(year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
