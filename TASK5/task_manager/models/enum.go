package models

type TaskStatus string

const (
	TaskDone TaskStatus = "done"
	TaskInProgress TaskStatus = "inprogress"
)
