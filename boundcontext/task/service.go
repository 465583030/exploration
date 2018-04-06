package task

import "github.com/nienie/exploration/boundcontext/usercontext"

//TaskService ...
type TaskService struct {}

//GetAvailableTasks ...
func (s *TaskService)GetAvailableTasks(ctx *usercontext.UserContext, completedTasks []*Task) []*Task {
    return nil
}