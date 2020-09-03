package storage

import (
	"context"
	"time"

	"github.com/scottshotgg/zeigarnik/pkg/types/dbtypes"
)

type Storage interface {
	ListReminders(ctx context.Context) ([]string, error)

	GetTTL(ctx context.Context, key string) (time.Duration, error)

	CreateReminder(ctx context.Context, r *dbtypes.Reminder) error
	GetReminder(ctx context.Context, key string) (*dbtypes.Reminder, error)
	UpdateReminder(ctx context.Context, r *dbtypes.Reminder) error
	DeleteReminder(ctx context.Context, key string) error
}
