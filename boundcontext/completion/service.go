package completion

import (
    "github.com/nienie/exploration/boundcontext/task"
    "github.com/nienie/exploration/boundcontext/execution"
    "github.com/nienie/exploration/boundcontext/usercontext"
)

//CompletionService ...
type CompletionService struct {}

//Complete ...
func (s *CompletionService)Complete(ctx *usercontext.UserContext, progress *execution.Progress) (bool, error) {
    return true, nil
}

//IsAlreadyCompleted ...
func (t *CompletionService)IsAlreadyCompleted(ctx *usercontext.UserContext, task *task.Task) bool {
    return false
}