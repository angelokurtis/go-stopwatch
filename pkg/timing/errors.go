package timing

// task_stop_error
type TaskStopError struct {
	err string
}

func (a *TaskStopError) Error() string {
	return a.err
}

func newNotRunningError() *TaskStopError {
	return &TaskStopError{err: "Can't stop StopWatch: it's not running"}
}
