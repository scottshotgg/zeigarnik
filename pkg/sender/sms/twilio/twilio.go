package twilio

import (
	"github.com/kevinburke/twilio-go"
	"github.com/scottshotgg/zeigarnik/pkg/sender"
	"github.com/scottshotgg/zeigarnik/pkg/storage/sql"
)

type TwilioSMS struct {
	from   string
	sid    string
	token  string
	client *twilio.Client
}

func New(f string, sid, token string) (sender.Sender, error) {
	// var err = sms.ValidateFrom(f)
	// if err != nil {
	// 	return nil, err
	// }

	// TODO: need some way to test the the client actually works
	return &TwilioSMS{
		from:   f,
		client: twilio.NewClient(sid, token, nil),
	}, nil
}

func (s *TwilioSMS) Send(r *sql.Reminder) error {
	var _, err = s.client.Messages.SendMessage(s.from, r.Recip, r.Msg, nil)
	return err
}
