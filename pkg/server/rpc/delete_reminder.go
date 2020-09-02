package rpc

import (
	"context"

	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
)

func (s *ReminderService) DeleteReminder(ctx context.Context, req *buffs.DeleteReminderReq) (*buffs.DeleteReminderRes, error) {
	var err = ctx.Err()
	if err != nil {
		return nil, err
	}

	err = s.store.DeleteReminder(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &buffs.DeleteReminderRes{}, nil
}
