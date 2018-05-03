package dvtask

type DvExecutionContext struct {
	Tasks []DvTask
}

func (context DvExecutionContext) initContext(tasks []DvTask) error {
	context.Tasks = tasks
	return nil
}

func (context DvExecutionContext) executeTasks() error {
	return nil
}
