package rpc

import (
	"context"
	"errors"

	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
)

func (s *ReminderService) UpdateReminder(context.Context, *buffs.UpdateReminderReq) (*buffs.UpdateReminderRes, error) {
	return nil, errors.New("not implemented")
}
