package sender

import "github.com/scottshotgg/zeigarnik/pkg/types/models"

type Sender interface {
	Send(r *models.Reminder) error
}
