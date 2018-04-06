package completion


import (
    "github.com/nienie/exploration/model/dao"
)

//CompletedTask ...
type UserCompletedTask struct {
    ID               int64
    UserID           int64
    TaskID           int64
    CompletedTaskDao *dao.CompletedTaskDao
}

//UserPrize ...
type UserPrize struct {
    ID               int64
    UserID           int64
    TaskID           int64
    UserPrize        *dao.UserPrizeDao
}