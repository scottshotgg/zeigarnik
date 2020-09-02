package mem

import (
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

func (s *Mem) ListReminders() ([]string, error) {
	return []string{}, nil
}

func (s *Mem) GetReminder(key string) (*dbtypes.Reminder, error) {
	var r, ok = s.store[key]
	if !ok {
		return nil, errors.New("not found")
	}

	return r, nil
}

func (s *Mem) GetTTL(key string) (time.Duration, error) {
	var r, err = s.GetReminder(key)
	if err != nil {
		return 0, err
	}

	return time.Duration(r.Moment - time.Now().Unix()), nil
}

func (s *Mem) CreateReminder(r *dbtypes.Reminder) error {
	s.store[r.ID] = r

	return nil
}

func (s *Mem) DeleteKey(key string) error {
	delete(s.store, key)

	return nil
}
