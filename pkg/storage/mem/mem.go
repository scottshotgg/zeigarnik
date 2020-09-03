package mem

import (
	"context"
	"errors"
	"time"

	"github.com/scottshotgg/zeigarnik/pkg/storage"
	"github.com/scottshotgg/zeigarnik/pkg/types/dbtypes"
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

func (s *Mem) ListReminders(_ context.Context) ([]string, error) {
	var keys []string

	for k := range s.store {
		keys = append(keys, k)
	}

	return keys, nil
}

func (s *Mem) GetReminder(_ context.Context, key string) (*dbtypes.Reminder, error) {
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

	return time.Duration(r.When - time.Now().Unix()), nil
}

func (s *Mem) CreateReminder(_ context.Context, r *dbtypes.Reminder) error {
	s.store[r.ID] = r

	return nil
}

func (s *Mem) UpdateReminder(_ context.Context, r *dbtypes.Reminder) error {
	// TODO: implement this
	return errors.New("not implemented")
}

func (s *Mem) DeleteReminder(_ context.Context, key string) error {
	delete(s.store, key)

	return nil
}
