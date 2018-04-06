package execution

import (
    "github.com/nienie/exploration/boundedcontext/task"
    "github.com/nienie/exploration/boundedcontext/usercontext"
)

//Factory ...
type Factory struct {}

//CreateProgress ...
func (f *Factory)CreateProgress(task *task.Task, ctx *usercontext.UserContext) (*Progress, error) {
    return nil, nil
}