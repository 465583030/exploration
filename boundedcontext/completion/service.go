package completion

import (
    "github.com/nienie/exploration/boundedcontext/task"
    "github.com/nienie/exploration/boundedcontext/usercontext"
)

//CompletionService 完成领域服务
type CompletionService struct {}

//Complete 完成任务后，所做的逻辑（记录用户完成某任务，写统计日志等工作）
func (s *CompletionService)Complete(ctx *usercontext.UserContext, task *task.Task) (bool, error) {
    return true, nil
}

//IsAlreadyCompleted 判断任务是否已经完成
func (t *CompletionService)IsAlreadyCompleted(ctx *usercontext.UserContext, task *task.Task) bool {
    return false
}