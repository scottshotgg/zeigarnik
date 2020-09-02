package rpc

import (
	"context"

	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
)

func (s *ReminderService) GetReminder(ctx context.Context, req *buffs.GetReminderByIDReq) (*buffs.GetReminderByIDRes, error) {
	var err = ctx.Err()
	if err != nil {
		return nil, err
	}

	r, err := s.store.GetReminder(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &buffs.GetReminderByIDRes{
		Reminder: &buffs.Reminder{
			Id:      r.ID,
			Created: r.Created,
			Message: r.Message,
			To:      r.To,
			Status:  r.Status,
			Moment:  r.Moment,
		},
	}, nil
}
