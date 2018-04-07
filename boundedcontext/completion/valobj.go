package completion

//CompletedTask ...
type UserCompletedTask struct {
    UserID           int64
    TaskID           int64
    CompletedTaskDao *UserCompletedTaskDao
}
