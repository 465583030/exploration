package completion

//UserCompletedTaskSet ...
type UserCompletedTaskSet struct {
    UserID          int64
    CompletedTasks  []*UserCompletedTask
}