package models

type TaskStatus string

const (
	TaskDone TaskStatus = "done"
	TaskInProgress TaskStatus = "inprogress"
	TaskPending TaskStatus = "pending"
)

func IsValidStatus(s TaskStatus) bool {
	switch s {
	case TaskDone, TaskInProgress, TaskPending:
		return true
	}
	return false
}