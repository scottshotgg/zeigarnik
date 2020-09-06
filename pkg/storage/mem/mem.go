package mem

import (
	"context"
	"errors"
	"time"

	"github.com/scottshotgg/zeigarnik/pkg/storage"
	"github.com/scottshotgg/zeigarnik/pkg/storage/sql"
)

type Mem struct {
	store map[string]*sql.Reminder
}

func New(kv map[string]*sql.Reminder) storage.Storage {
	if kv == nil {
		kv = map[string]*sql.Reminder{}
	}

	return &Mem{
		store: kv,
	}
}

func (s *Mem) ListReminders(_ context.Context) ([]string, error) {
	var keys []string

	for k := range s.store {
		keys = append(keys, k)
	}

	return keys, nil
}

func (s *Mem) GetReminder(_ context.Context, key string) (*sql.Reminder, error) {
	var r, ok = s.store[key]
	if !ok {
		return nil, errors.New("not found")
	}

	return r, nil
}

func (s *Mem) GetTTL(_ context.Context, key string) (time.Duration, error) {
	var r, err = s.GetReminder(nil, key)
	if err != nil {
		return 0, err
	}

	return time.Duration(r.Atwhen - time.Now().Unix()), nil
}

func (s *Mem) CreateReminder(_ context.Context, r *sql.Reminder) error {
	s.store[r.ID.String()] = r

	return nil
}

func (s *Mem) UpdateReminder(_ context.Context, r *sql.Reminder) error {
	// TODO: implement this
	return errors.New("not implemented")
}

func (s *Mem) DeleteReminder(_ context.Context, key string) error {
	delete(s.store, key)

	return nil
}
