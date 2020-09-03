package rpc

import (
	"context"

	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
)

func (s *ReminderService) DeleteReminder(ctx context.Context, req *reminder.DeleteReminderReq) (*reminder.DeleteReminderRes, error) {
	var err = ctx.Err()
	if err != nil {
		return nil, err
	}

	err = s.store.DeleteReminder(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &reminder.DeleteReminderRes{}, nil
}
