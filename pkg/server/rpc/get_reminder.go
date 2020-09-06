package rpc

import (
	"context"

	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
	"github.com/scottshotgg/zeigarnik/pkg/types/transform"
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

	return &reminder.GetReminderByIDRes{
		Reminder: transform.ReminderToPbPtr(r),
	}, nil
}
