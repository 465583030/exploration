package prize

//Prize ...
type Prize struct {
    PrizeDao *PrizeDao
}

//TaskPrize ...
type TaskPrize struct {
    TaskID      int64
    Prizes      []*Prize
    TaskPrizeDao *TaskPrizeDao
}

//UserPrize ...
type UserPrize struct {
    UserPrize        *UserPrizeDao
}