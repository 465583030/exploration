package task

import "github.com/nienie/exploration/boundedcontext/usercontext"

//TaskService 任务领域服务
type TaskService struct {}

//GetAvailableTasks 获取用户可以执行的任务。
func (s *TaskService)GetAvailableTasks(ctx *usercontext.UserContext, completedTasks []*Task) []*Task {
    return nil
}