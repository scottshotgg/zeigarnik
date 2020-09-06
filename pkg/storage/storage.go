package storage

import (
	"context"
	"time"

	"github.com/scottshotgg/zeigarnik/pkg/storage/sql"
)

type Storage interface {
	ListReminders(ctx context.Context) ([]string, error)

	GetTTL(ctx context.Context, key string) (time.Duration, error)

	CreateReminder(ctx context.Context, r *sql.Reminder) error
	GetReminder(ctx context.Context, key string) (*sql.Reminder, error)
	UpdateReminder(ctx context.Context, r *sql.Reminder) error
	DeleteReminder(ctx context.Context, key string) error
}
