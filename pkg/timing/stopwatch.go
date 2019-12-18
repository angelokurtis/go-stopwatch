package timing

import (
	"fmt"
	"strings"
	"time"
)

type StopWatch struct {
	id              string
	taskList        []TaskInfo
	startTime       time.Time
	currentTaskName string
	taskCount       int
	totalTime       time.Duration
	lastTaskInfo    *TaskInfo
}

func (sw *StopWatch) Start(taskName string) error {
	if sw.isRunning() {
		return newAlreadyRunningError()
	}
	sw.currentTaskName = taskName
	sw.startTime = time.Now()
	return nil
}

func (sw *StopWatch) Stop() error {
	if !sw.isRunning() {
		return newNotRunningError()
	}
	lastTime := time.Now().Sub(sw.startTime)
	sw.totalTime = sw.totalTime + lastTime
	sw.lastTaskInfo = newTaskInfo(sw.currentTaskName, lastTime)
	sw.taskList = append(sw.taskList, *sw.lastTaskInfo)
	sw.taskCount++
	sw.currentTaskName = ""
	return nil
}

func (sw *StopWatch) PrettyPrint() string {
	var sb strings.Builder
	for _, t := range sw.taskList {
		p := t.time.Seconds() / sw.totalTime.Seconds()
		line := fmt.Sprintf("%s  %f  %s\n", t.time.String(), p, t.taskName)
		sb.WriteString(line)
	}
	return sb.String()
}

func (sw *StopWatch) isRunning() bool {
	return sw.currentTaskName != ""
}

func (sw *StopWatch) shortSummary() string {
	return "StopWatch '" + sw.id + "': running time = " + sw.totalTime.String()
}
