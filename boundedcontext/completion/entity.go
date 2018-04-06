package completion

//CompletedTask ...
type UserCompletedTask struct {
    ID               int64
    UserID           int64
    TaskID           int64
    CompletedTaskDao *CompletedTaskDao
}

//UserPrize ...
type UserPrize struct {
    ID               int64
    UserID           int64
    TaskID           int64
    UserPrize        *UserPrizeDao
}