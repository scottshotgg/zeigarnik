package storage

import (
	"time"

	"github.com/scottshotgg/zeigarnik/pkg/dbtypes"
)

type Storage interface {
	ListReminders() ([]string, error)

	GetReminder(key string) (*dbtypes.Reminder, error)
	GetTTL(key string) (time.Duration, error)

	CreateReminder(r *dbtypes.Reminder) error
	DeleteKey(key string) error
}
