package task

//Task ...
type Task struct {
    TaskID      int64
    TaskDao     *TaskDao
    Condition   *Condition
}