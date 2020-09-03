package naive

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/scottshotgg/zeigarnik/pkg/sender"
	"github.com/scottshotgg/zeigarnik/pkg/storage"
	"github.com/scottshotgg/zeigarnik/pkg/types/dbtypes"
	"github.com/scottshotgg/zeigarnik/pkg/types/models"
)

type (
	Naive struct {
		store  storage.Storage
		sender sender.Sender
		ch     chan string
	}
)

func New(store storage.Storage) *Naive {
	return &Naive{
		store: store,
	}
}
func (s *Naive) Start() {
	go s.Scrape()

	// Start the workers
	for i := 0; i < 10; i++ {
		go s.worker()
	}
}

func (s *Naive) Scrape() {
	for range time.NewTicker(1 * time.Minute).C {
		var keys, err = s.store.ListReminders(context.Background())
		if err != nil {
			log.Fatalln("err:", err)

			// TODO: need to do something to decode the error and act
			continue
		}

		for _, key := range keys {
			s.ch <- key
		}
	}
}

func (s *Naive) worker() {
	var err error

	for key := range s.ch {
		err = s.process(key)
		if err != nil {
			// TODO: do something here

			log.Fatalln("err:", err)
		}
	}
}

/*
 - Get all the keys
 - Loop through the keys
 - Set time.AfterFunc for each one under 5 minutes
*/
func (s *Naive) process(key string) error {
	var (
		ctx = context.Background()

		// Grab the TTL of the given key
		ttl, err = s.store.GetTTL(ctx, key)
	)

	if err != nil {
		log.Fatalln("err:", err)
	}

	// If TTL is over 5 minutes, ignore it for now
	if ttl > 5*time.Minute {
		fmt.Println("Skipping ...")
		return nil
	}

	// Get the key
	r, err := s.store.GetReminder(ctx, key)
	if err != nil {
		log.Fatalln("err:", err)
		return err
	}

	// If its already been queued then skip it
	// TODO: put this in another Redis hash set or w/e
	switch r.Status {
	case dbtypes.CREATED:
		// Set a timer to fire the send
		s.remind(&models.Reminder{
			ID:      r.ID,
			Created: r.Created,
			When:    r.When,
			Message: r.Message,
			Status:  r.Status,
			To:      r.To,
		})

	case
		dbtypes.QUEUED,
		dbtypes.FIRED,
		dbtypes.MISSED:

		fmt.Println("Already " + r.Status)

	default:
		return errors.New("invalid message status: " + r.Status)
	}

	return nil
}

// TODO: calculating the TTL needs to be more calculated, current time, etc
func (s *Naive) remind(r *models.Reminder) {
	// TODO: use the context here
	time.AfterFunc(r.CalcTTL(), func() {
		var err = s.sender.Send(r)
		if err != nil {
			// TODO: need to do something here
			log.Fatalln("err:", err)
		}

		// TODO: we can probably optimize this by batching them
		// Delete the key after we are done with it
		err = s.store.DeleteReminder(context.Background(), r.ID)
		if err != nil {
			log.Fatalln("err:", err)
		}
	})
}
