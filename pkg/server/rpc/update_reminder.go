package rpc

import (
	"context"

	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
	"github.com/scottshotgg/zeigarnik/pkg/dbtypes"
)

func (s *ReminderService) UpdateReminder(ctx context.Context, req *buffs.UpdateReminderReq) (*buffs.UpdateReminderRes, error) {
	var err = ctx.Err()
	if err != nil {
		return nil, err
	}

	err = s.store.UpdateReminder(ctx, dbtypes.ToDB(req.GetReminder()))
	if err != nil {
		return nil, err
	}

	return &buffs.UpdateReminderRes{
		Reminder: req.GetReminder(),
	}, nil
}
