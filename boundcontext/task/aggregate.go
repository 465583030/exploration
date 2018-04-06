package task

import (
    "github.com/nienie/exploration/model/dao"
)

//Task ...
type Task struct {
    TaskID      int64
    TaskDao     *dao.TaskDao
    Condition   *Condition
}