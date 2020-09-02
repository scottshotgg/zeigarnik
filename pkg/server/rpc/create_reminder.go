package rpc

import (
	"context"
	"errors"

	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
)

func (s *ReminderService) CreateReminder(context.Context, *buffs.CreateReminderReq) (*buffs.CreateReminderRes, error) {
	return nil, errors.New("not implemented")
}
