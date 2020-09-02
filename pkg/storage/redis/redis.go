package redis

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/kelindar/binary"
	"github.com/scottshotgg/zeigarnik/pkg/dbtypes"
	"github.com/scottshotgg/zeigarnik/pkg/storage"
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

func (s *Redis) ListReminders() ([]string, error) {
	return s.client.Keys(everything).Result()
}

func (s *Redis) GetReminder(key string) (*dbtypes.Reminder, error) {
	var contents, err = s.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}

	// Unmarshal the contents from the binary payload
	var r dbtypes.Reminder
	err = binary.Unmarshal(contents, &r)
	if err != nil {
		log.Fatalln("err:", err)
	}

	return &r, nil
}

func (s *Redis) GetTTL(key string) (time.Duration, error) {
	// Grab the TTL of the given key
	return s.client.TTL(key).Result()
}

func (s *Redis) CreateReminder(r *dbtypes.Reminder) error {
	// If not then insert it into Redis
	blob, err := binary.Marshal(r)
	if err != nil {
		log.Fatalln("err:", err)
	}

	_, err = s.client.Set(r.ID, blob, time.Duration(r.Moment-time.Now().Unix())).Result()

	return err
}

func (s *Redis) DeleteKey(key string) error {
	// TODO: what to do with count
	var _, err = s.client.Del(key).Result()

	// if err != nil {
	// 	return err
	// }

	return err
}
