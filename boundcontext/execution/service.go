package execution

import (
    "github.com/nienie/exploration/boundcontext/task"
    "github.com/nienie/exploration/boundcontext/usercontext"
)

//ExecutionService ...
type ExecutionService struct {}

//UpdateProgress ...
func (s *ExecutionService)UpdateProgress(ctx usercontext.UserContext, task *task.Task) error {
    return nil
}