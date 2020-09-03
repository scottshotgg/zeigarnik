package sender

import (
	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
)

type Sender interface {
	Send(r *reminder.Reminder) error
}
