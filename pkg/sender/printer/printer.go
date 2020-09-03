package printer

import (
	"log"

	"github.com/scottshotgg/zeigarnik/pkg/sender"
	"github.com/scottshotgg/zeigarnik/pkg/types/models"
)

// Printer implements Sender but is just a mock/test to just print
type Printer struct{}

func New() (sender.Sender, error) {
	return &Printer{}, nil
}

// Send is implemented very simply using time.AfterFunc
func (p *Printer) Send(r *models.Reminder) error {
	log.Println("MESSAGE:", r.Message)

	return nil
}
