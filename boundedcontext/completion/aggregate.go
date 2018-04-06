package completion

//UserCompletedTaskRepo ...
type UserCompletedTaskRepo struct {
    UserID          int64
    CompletedTasks  []*UserCompletedTask
}