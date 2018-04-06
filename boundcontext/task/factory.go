package task

//Factory ...
type Factory struct {}

//GetTaskByID ...
func (f *Factory)GetTaskByID(taskID int64) *Task{
    return nil
}

//CreateCondition ...
func (f *Factory)CreateCondition(bytes []byte) (*Condition, error) {
    return nil, nil
}