package model

type Status int

const (
	// StatusAny 任何状态
	StatusAny Status = iota
	// StatusPending 代办
	StatusPending
	// StatusCompleted 已完成
	StatusCompleted
)
