package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/kelindar/binary"
	"github.com/scottshotgg/zeigarnik/pkg/storage"
	"github.com/scottshotgg/zeigarnik/pkg/storage/sql"
)

const (
	everything = "*"
)

type Redis struct {
	client *redis.Client
}

func New(uri string) (storage.Storage, error) {
	var (
		// TODO: options
		client = redis.NewClient(&redis.Options{
			Addr:            uri,
			MaxRetries:      3,
			MinRetryBackoff: 1 * time.Second,
			MaxRetryBackoff: 10 * time.Second,
			// Password: "", // no password set
			// DB:       0,  // use default DB
		})

		err = client.Ping().Err()
	)

	if err != nil {
		return nil, err
	}

	return &Redis{
		client: client,
	}, nil
}

func (s *Redis) ListReminders(ctx context.Context) ([]string, error) {
	return s.client.WithContext(ctx).Keys(everything).Result()
}

func (s *Redis) GetReminder(ctx context.Context, key string) (*sql.Reminder, error) {
	var contents, err = s.client.WithContext(ctx).Get(key).Bytes()
	if err != nil {
		return nil, err
	}

	var r sql.Reminder

	// Unmarshal the contents from the binary payload
	err = binary.Unmarshal(contents, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (s *Redis) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	// Grab the TTL of the given key
	return s.client.WithContext(ctx).TTL(key).Result()
}

func (s *Redis) CreateReminder(ctx context.Context, r *sql.Reminder) error {
	// If not then insert it into Redis
	blob, err := binary.Marshal(r)
	if err != nil {
		return err
	}

	_, err = s.client.WithContext(ctx).Set(r.ID.String(), blob, time.Duration(r.Atwhen-time.Now().Unix())).Result()

	return err
}

func (s *Redis) UpdateReminder(ctx context.Context, r *sql.Reminder) error {
	// // If not then insert it into Redis
	// blob, err := binary.Marshal(r)
	// if err != nil {
	// 	log.Fatalln("err:", err)
	// }

	// _, err = s.client.Set(r.ID, blob, time.Duration(r.Moment-time.Now().Unix())).Result()

	// return err

	return errors.New("not implemented")
}

func (s *Redis) DeleteReminder(ctx context.Context, key string) error {
	// TODO: what to do with count
	var _, err = s.client.WithContext(ctx).Del(key).Result()

	// if err != nil {
	// 	return err
	// }

	return err
}
