package rpc

import (
	"context"
	"errors"

	buffs "github.com/scottshotgg/zeigarnik/pkg/buffs"
)

func (s *ReminderService) GetReminder(context.Context, *buffs.GetReminderByIDReq) (*buffs.GetReminderByIDRes, error) {
	return nil, errors.New("not implemented")
}
