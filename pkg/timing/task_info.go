package timing

import "time"

type TaskInfo struct {
	taskName string
	time     time.Duration
}

func newTaskInfo(taskName string, time time.Duration) *TaskInfo {
	return &TaskInfo{taskName: taskName, time: time}
}
