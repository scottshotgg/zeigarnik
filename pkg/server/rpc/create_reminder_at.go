package rpc

import (
	"context"

	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
	"github.com/scottshotgg/zeigarnik/pkg/types/transform"
)

func (s *ReminderService) CreateReminder(ctx context.Context, req *reminder.CreateReminderReq) (*reminder.CreateReminderRes, error) {
	var err = ctx.Err()
	if err != nil {
		return nil, err
	}

	err = s.store.CreateReminder(ctx, transform.PbToReminderPtr(req.GetReminder()))
	if err != nil {
		return nil, err
	}

	return &reminder.CreateReminderRes{
		Reminder: req.GetReminder(),
	}, nil
}
