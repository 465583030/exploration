package execution

import (
    "github.com/nienie/exploration/boundcontext/task"
    "github.com/nienie/exploration/boundcontext/usercontext"
)

//Factory ...
type Factory struct {}

//CreateProgress ...
func (f *Factory)CreateProgress(task *task.Task, ctx *usercontext.UserContext) (*Progress, error) {
    return nil, nil
}