package rpc

import (
	"context"

	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
	"github.com/scottshotgg/zeigarnik/pkg/types/transform"
)

func (s *ReminderService) UpdateReminder(ctx context.Context, req *reminder.UpdateReminderReq) (*reminder.UpdateReminderRes, error) {
	var err = ctx.Err()
	if err != nil {
		return nil, err
	}

	err = s.store.UpdateReminder(ctx, transform.PbToReminderPtr(req.GetReminder()))
	if err != nil {
		return nil, err
	}

	return &reminder.UpdateReminderRes{
		Reminder: req.GetReminder(),
	}, nil
}
