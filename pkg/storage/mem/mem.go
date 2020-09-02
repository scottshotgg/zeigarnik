package mem

import (
	"context"
	"errors"
	"time"

	"github.com/scottshotgg/zeigarnik/pkg/dbtypes"
	"github.com/scottshotgg/zeigarnik/pkg/storage"
)

type Mem struct {
	store map[string]*dbtypes.Reminder
}

func New(kv map[string]*dbtypes.Reminder) storage.Storage {
	if kv == nil {
		kv = map[string]*dbtypes.Reminder{}
	}

	return &Mem{
		store: kv,
	}
}

func (s *Mem) ListReminders(ctx context.Context) ([]string, error) {
	return []string{}, nil
}

func (s *Mem) GetReminder(ctx context.Context, key string) (*dbtypes.Reminder, error) {
	var r, ok = s.store[key]
	if !ok {
		return nil, errors.New("not found")
	}

	return r, nil
}

func (s *Mem) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	var r, err = s.GetReminder(ctx, key)
	if err != nil {
		return 0, err
	}

	return time.Duration(r.Moment - time.Now().Unix()), nil
}

func (s *Mem) CreateReminder(ctx context.Context, r *dbtypes.Reminder) error {
	s.store[r.ID] = r

	return nil
}

func (s *Mem) UpdateReminder(ctx context.Context, r *dbtypes.Reminder) error {
	return errors.New("not implemented")
}

func (s *Mem) DeleteReminder(ctx context.Context, key string) error {
	delete(s.store, key)

	return nil
}
