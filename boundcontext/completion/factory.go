package completion

import (
    "github.com/nienie/exploration/boundcontext/task"
    "github.com/nienie/exploration/boundcontext/usercontext"
)

//Factory ...
type Factory struct {}

//CreateUserCompletedTasks ...
func (f *Factory)CreateUserCompletedTasks(ctx *usercontext.UserContext) (*UserCompletedTaskRepo, error) {
    return nil, nil
}

//GetUserCompletedTasks ...
func (f *Factory)GetUserCompletedTasks(ctx *usercontext.UserContext) ([]*task.Task, error) {
    return nil, nil
}

//CreateCompletedTasks ...
func (f *Factory)CreateCompletedTasks(ctx *usercontext.UserContext, task *task.Task) (*UserCompletedTask, error) {
    return nil, nil
}