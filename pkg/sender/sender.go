package sender

import "github.com/scottshotgg/zeigarnik/pkg/storage/sql"

type Sender interface {
	Send(r *sql.Reminder) error
}
