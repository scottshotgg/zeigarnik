package rpc

import (
	"context"

	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
)

func (s *ReminderService) GetReminder(ctx context.Context, req *reminder.GetReminderByIDReq) (*reminder.GetReminderByIDRes, error) {
	var err = ctx.Err()
	if err != nil {
		return nil, err
	}

	r, err := s.store.GetReminder(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	var status = reminder.ReminderStatus(reminder.ReminderStatus_value[r.Status])

	return &reminder.GetReminderByIDRes{
		Reminder: &reminder.Reminder{
			Id:      r.ID,
			Created: r.Created,
			Message: r.Message,
			To:      r.To,
			Status:  status,
			When:    r.When,
		},
	}, nil
}
