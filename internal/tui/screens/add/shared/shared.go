package shared

import (
	"time"

	"github.com/DerBaumann/timelog/internal/store"
)

type FormData struct {
	Project     string
	StartTime   store.Minutes
	EndTime     store.Minutes
	Duration    time.Duration
	Description string
}
