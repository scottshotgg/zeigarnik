package helpers

import (
	"github.com/google/uuid"
	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
)

func UuidUUIDToString(id uuid.UUID) string {
	return id.String()
}

func StringToUuidUUID(id string) uuid.UUID {
	var idd, err = uuid.Parse(id)
	if err != nil {
		return uuid.Nil
	}

	return idd
}

func StatusToString(r reminder.ReminderStatus) string {
	return r.String()
}

func TypeToString(r reminder.ReminderType) string {
	return r.String()
}

func StringToStatus(r string) reminder.ReminderStatus {
	return reminder.ReminderStatus(reminder.ReminderStatus_value[r])
}

func StringToType(r string) reminder.ReminderType {
	return reminder.ReminderType(reminder.ReminderType_value[r])
}
