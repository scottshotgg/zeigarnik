package rpc

import (
	"context"
	"errors"

	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
)

func (s *ReminderService) DeleteReminder(context.Context, *buffs.DeleteReminderReq) (*buffs.DeleteReminderRes, error) {
	return nil, errors.New("not implemented")
}
