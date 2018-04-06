package execution

import (
    "github.com/nienie/exploration/boundedcontext/task"
    "github.com/nienie/exploration/boundedcontext/usercontext"
)

//ExecutionService 执行领域服务
type ExecutionService struct {}

//UpdateProgress 更新任务的完成进度
func (s *ExecutionService)UpdateProgress(ctx usercontext.UserContext, task *task.Task) (*Progress, error) {
    return nil, nil
}

//GetProgress 获取任务的执行进度
func (s *ExecutionService)GetProgress(ctx usercontext.UserContext, task *task.Task) (*Progress, error) {
    return nil, nil
}

//IsProgressCompleted 判断任务进度是否是完成
func (s *ExecutionService)IsProgressCompleted(progress *Progress) bool {
    return false
}