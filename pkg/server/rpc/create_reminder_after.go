package rpc

// import (
// 	"context"

// 	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
// 	"github.com/scottshotgg/zeigarnik/pkg/dbtypes"
// )

// func (s *ReminderService) CreateReminder(ctx context.Context, req *buffs.CreateReminderReq) (*buffs.CreateReminderRes, error) {
// 	var err = ctx.Err()
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = s.store.CreateReminder(ctx, dbtypes.ToDB(req.GetReminder()))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &buffs.CreateReminderRes{
// 		Reminder: req.GetReminder(),
// 	}, nil
// }
